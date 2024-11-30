package usersvc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/user"
	authsvc "github.com/ruhollahh/go-ecom/internal/service/auth"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

var (
	ErrNotFound          = errors.New("user not found")
	ErrUniquePhoneNumber = errors.New("phone number is not unique")
	ErrWrongCredentials  = errors.New("wrong credentials")
)

type Storer interface {
	Create(ctx context.Context, usr user.User) error
	QueryByID(ctx context.Context, userID uuid.UUID) (user.User, error)
	QueryByPhoneNumber(ctx context.Context, phoneNumber string) (user.User, error)
}

type AuthSvc interface {
	CreateAccessToken(userID uuid.UUID) (string, error)
	CreateRefreshToken(userID uuid.UUID) (string, error)
	ParseRefreshToken(refreshToken string) (authsvc.Claims, error)
}

type Service struct {
	storer  Storer
	authSvc AuthSvc
	log     *logger.Logger
}

func New(log *logger.Logger, authSvc AuthSvc, storer Storer) *Service {
	return &Service{
		storer:  storer,
		authSvc: authSvc,
		log:     log,
	}
}

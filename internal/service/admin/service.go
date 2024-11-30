package adminsvc

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/admin"
	authsvc "github.com/ruhollahh/go-ecom/internal/service/auth"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

var (
	ErrNotFound         = errors.New("admin not found")
	ErrWrongCredentials = errors.New("wrong credentials")
)

type Storer interface {
	QueryByPhoneNumber(ctx context.Context, phoneNumber string) (admin.Admin, error)
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

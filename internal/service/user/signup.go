package usersvc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type SignupReq struct {
	Name        string
	PhoneNumber string
	Password    string
}

type SignupResp struct {
	ID uuid.UUID
}

func (s *Service) Signup(ctx context.Context, req SignupReq) (SignupResp, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return SignupResp{}, fmt.Errorf("generateFromPassword: %w", err)
	}

	now := time.Now()

	usr := user.User{
		ID:             uuid.New(),
		Name:           req.Name,
		PhoneNumber:    req.PhoneNumber,
		HashedPassword: hash,
		Active:         true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if err = s.storer.Create(ctx, usr); err != nil {

		return SignupResp{}, fmt.Errorf("create: %w", err)
	}

	return SignupResp{
		ID: usr.ID,
	}, nil
}

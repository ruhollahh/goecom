package usersvc

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	PhoneNumber string
	Password    string
}

type LoginResp struct {
	AccessToken  string
	RefreshToken string
}

func (s *Service) Login(ctx context.Context, req LoginReq) (LoginResp, error) {
	user, err := s.storer.QueryByPhoneNumber(ctx, req.PhoneNumber)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			return LoginResp{}, ErrWrongCredentials
		}
		return LoginResp{}, fmt.Errorf("query: phone number[%s]: %w", req.PhoneNumber, err)
	}

	err = bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return LoginResp{}, ErrWrongCredentials
		}

		return LoginResp{}, fmt.Errorf("compareHashAndPassword: %w", err)
	}

	accessToken, err := s.authSvc.CreateAccessToken(user.ID)
	if err != nil {
		return LoginResp{}, fmt.Errorf("createAccessToken: %w", err)
	}

	refreshToken, err := s.authSvc.CreateRefreshToken(user.ID)
	if err != nil {
		return LoginResp{}, fmt.Errorf("createRefreshToken: %w", err)
	}

	return LoginResp{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

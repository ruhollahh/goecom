package authsvc

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

var (
	ErrAuthFailed          = errors.New("authentication failed")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

type Config struct {
	SignKey               []byte
	AccessExpirationTime  time.Duration
	RefreshExpirationTime time.Duration
	AccessSubject         string
	RefreshSubject        string
}

type Claims struct {
	jwt.RegisteredClaims
	UserID uuid.UUID `json:"user_id"`
}

type Service struct {
	cfg Config
}

func NewService(cfg Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

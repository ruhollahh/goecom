package authsvc

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"time"
)

func (s Service) CreateAccessToken(userID uuid.UUID) (string, error) {
	return s.createToken(userID, s.cfg.AccessSubject, s.cfg.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(userID uuid.UUID) (string, error) {
	return s.createToken(userID, s.cfg.RefreshSubject, s.cfg.RefreshExpirationTime)
}

func (s Service) createToken(userID uuid.UUID, subject string, expireDuration time.Duration) (string, error) {
	c := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
			Subject:   subject,
		},
		UserID: userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(s.cfg.SignKey)
}

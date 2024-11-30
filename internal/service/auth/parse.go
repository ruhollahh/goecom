package authsvc

import (
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

func (s Service) Authenticate(bearerToken string) (Claims, error) {
	parts := strings.Split(bearerToken, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return Claims{}, ErrAuthFailed
	}

	return s.parseToken(parts[1])
}

func (s Service) ParseRefreshToken(refreshToken string) (Claims, error) {
	claims, err := s.parseToken(refreshToken)
	if err != nil {
		return Claims{}, err
	}
	if claims.Subject != s.cfg.RefreshSubject {
		return Claims{}, ErrInvalidRefreshToken
	}

	return claims, nil
}

func (s Service) parseToken(token string) (Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return Claims{}, ErrAuthFailed
		}

		return s.cfg.SignKey, nil
	})
	if err != nil {
		return Claims{}, ErrAuthFailed
	}
	if !parsedToken.Valid {
		return Claims{}, ErrAuthFailed
	}

	return parsedToken.Claims.(Claims), nil
}

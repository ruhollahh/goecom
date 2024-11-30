package claim

import (
	"context"
	"errors"
	authsvc "github.com/ruhollahh/go-ecom/internal/service/auth"
)

type ctxKey int

const claimKey ctxKey = 1

// SetClaims stores the claims in the context.
func SetClaims(ctx context.Context, claims authsvc.Claims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) (authsvc.Claims, error) {
	v, ok := ctx.Value(claimKey).(authsvc.Claims)
	if !ok {
		return authsvc.Claims{}, errors.New("claims not found in context")
	}
	return v, nil
}

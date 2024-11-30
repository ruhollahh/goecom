package usersvc

import (
	"context"
	"fmt"
	"github.com/ruhollahh/go-ecom/internal/entity/user"
)

func (s *Service) GetByPhoneNumber(ctx context.Context, phoneNumber string) (user.User, error) {
	usr, err := s.storer.QueryByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		return user.User{}, fmt.Errorf("query: phone number[%s]: %w", phoneNumber, err)
	}

	return usr, nil
}

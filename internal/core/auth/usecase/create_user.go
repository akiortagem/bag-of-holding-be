package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
)

type CreateUserUsecase struct {
	Service domain.AuthService
}

func (u *CreateUserUsecase) CreateUser(ctx context.Context, payload []byte) (domain.UserCreateResponse, error) {
	var params domain.UserCreateParams
	if err := json.Unmarshal(payload, &params); err != nil {
		return domain.UserCreateResponse{}, fmt.Errorf("decode create user params: %w", err)
	}

	hashedPassword, err := HashPassword(params.Password)
	if err != nil {
		return domain.UserCreateResponse{}, fmt.Errorf("hash password: %w", err)
	}
	params.Password = hashedPassword

	return u.Service.CreateUser(ctx, params)
}

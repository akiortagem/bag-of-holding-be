package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
)

type LoginUserUsecase struct {
	Service domain.AuthService
}

type LoginFailedError struct {
}

func (m *LoginFailedError) Error() string {
	return "Login failed"
}

func (u *LoginUserUsecase) LoginUser(ctx context.Context, payload []byte, cfg config.ApiConfig) (domain.UserLoginResponse, error) {
	var params domain.UserLoginParams
	if err := json.Unmarshal(payload, &params); err != nil {
		return domain.UserLoginResponse{}, fmt.Errorf("decode create user login: %w", err)
	}

	dbUser, err := u.Service.GetUserByEmail(ctx, domain.UserGetByEmailParams{Email: params.Email})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserLoginResponse{}, &LoginFailedError{}
		}
		return domain.UserLoginResponse{}, err
	}

	ok, err := CheckPasswordHash(params.Password, dbUser.PasswordHash)
	if err != nil {
		return domain.UserLoginResponse{}, fmt.Errorf("verify password: %w", err)
	}
	if !ok {
		return domain.UserLoginResponse{}, &LoginFailedError{}
	}

	token, err := MakeJWT(dbUser.ID, time.Duration(3600)*time.Second, cfg.JWTIssuer, cfg.Secret)

	if err != nil {
		return domain.UserLoginResponse{}, err
	}

	var rtVal string

	if rtVal, err = MakeRefreshToken(); err != nil {
		return domain.UserLoginResponse{}, err
	}

	if _, err := u.Service.CreateRefreshToken(ctx, domain.RefreshTokenCreateParams{
		Token:     rtVal,
		UserID:    dbUser.ID,
		ExpiresAt: time.Now().Add(time.Duration(3600) * time.Second),
	}); err != nil {
		return domain.UserLoginResponse{}, err
	}

	return domain.UserLoginResponse{
		Email:        dbUser.Email,
		Token:        token,
		RefreshToken: rtVal,
	}, nil

}

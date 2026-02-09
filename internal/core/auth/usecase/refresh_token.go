package usecase

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
)

type RefreshTokenUsecase struct {
	Service domain.AuthService
}

type RefreshTokenInvalidError struct{}

func (m *RefreshTokenInvalidError) Error() string {
	return "Refresh token invalid"
}

func (u *RefreshTokenUsecase) RefreshToken(ctx context.Context, refreshToken string, cfg config.ApiConfig) (domain.UserLoginResponse, error) {
	if refreshToken == "" {
		return domain.UserLoginResponse{}, &RefreshTokenInvalidError{}
	}

	dbToken, err := u.Service.GetRefreshToken(ctx, domain.RefreshTokenGetParams{Token: refreshToken})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserLoginResponse{}, &RefreshTokenInvalidError{}
		}
		return domain.UserLoginResponse{}, err
	}

	if dbToken.RevokedAt != nil {
		return domain.UserLoginResponse{}, &RefreshTokenInvalidError{}
	}

	if time.Now().After(dbToken.ExpiresAt) {
		return domain.UserLoginResponse{}, &RefreshTokenInvalidError{}
	}

	token, err := MakeJWT(dbToken.UserID, time.Duration(3600)*time.Second, cfg.JWTIssuer, cfg.Secret)
	if err != nil {
		return domain.UserLoginResponse{}, err
	}

	newRefreshToken, err := MakeRefreshToken()
	if err != nil {
		return domain.UserLoginResponse{}, err
	}

	if _, err := u.Service.CreateRefreshToken(ctx, domain.RefreshTokenCreateParams{
		Token:     newRefreshToken,
		UserID:    dbToken.UserID,
		ExpiresAt: time.Now().Add(time.Duration(3600) * time.Second),
	}); err != nil {
		return domain.UserLoginResponse{}, err
	}

	if err := u.Service.RevokeRefreshToken(ctx, domain.RefreshTokenRevokeParams{Token: refreshToken}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.UserLoginResponse{}, &RefreshTokenInvalidError{}
		}
		return domain.UserLoginResponse{}, err
	}

	return domain.UserLoginResponse{
		Token:        token,
		RefreshToken: newRefreshToken,
	}, nil
}

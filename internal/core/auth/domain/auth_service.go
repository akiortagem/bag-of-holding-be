package domain

import "context"

type AuthService interface {
	CreateUser(ctx context.Context, newUser UserCreateParams) (UserCreateResponse, error)
	CreateRefreshToken(ctx context.Context, params RefreshTokenCreateParams) (RefreshToken, error)
	GetRefreshToken(ctx context.Context, params RefreshTokenGetParams) (RefreshToken, error)
	RevokeRefreshToken(ctx context.Context, params RefreshTokenRevokeParams) error
	GetUserByEmail(ctx context.Context, params UserGetByEmailParams) (User, error)
}

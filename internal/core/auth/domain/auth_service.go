package domain

import "context"

type AuthService interface {
	CreateUser(ctx context.Context, newUser UserCreateParams) (UserCreateResponse, error)
	CreateRefreshToken(ctx context.Context, params RefreshTokenCreateParams) (RefreshToken, error)
	GetUserByEmail(ctx context.Context, params UserGetByEmailParams) (User, error)
}

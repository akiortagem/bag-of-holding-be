package domain

import "context"

type AuthService interface {
	CreateUser(ctx context.Context, newUser UserCreateParams) (UserCreateResponse, error)
}

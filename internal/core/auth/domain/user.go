package domain

import "time"

type UserCreateParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreateResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type UserLoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Email        string `json:"email"`
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserGetByEmailParams struct {
	Email string
}

type User struct {
	ID              int64
	Email           string
	PasswordHash    string
	IsActive        bool
	EmailVerifiedAt *time.Time
	LastLoginAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

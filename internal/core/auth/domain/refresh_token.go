package domain

import "time"

type RefreshTokenCreateParams struct {
	Token     string
	UserID    int64
	ExpiresAt time.Time
}

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
	ExpiresAt time.Time
	RevokedAt *time.Time
}


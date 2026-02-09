package domain

import "time"

type RefreshTokenCreateParams struct {
	Token     string
	UserID    int64
	ExpiresAt time.Time
}

type RefreshTokenGetParams struct {
	Token string
}

type RefreshTokenRevokeParams struct {
	Token string
}

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
	ExpiresAt time.Time
	RevokedAt *time.Time
}

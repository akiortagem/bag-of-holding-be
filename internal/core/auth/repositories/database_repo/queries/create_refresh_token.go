package queries

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DatabaseRefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type DatabaseRefreshTokenCreateParams struct {
	Token     string
	UserID    int64
	ExpiresAt time.Time
}

const createRefreshTokenQuery = `
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at)
VALUES ($1, NOW(), NOW(), $2, $3)
RETURNING token, created_at, updated_at, user_id, expires_at, revoked_at;
`

func CreateRefreshTokenQuery(ctx context.Context, db *sql.DB, params DatabaseRefreshTokenCreateParams) (DatabaseRefreshToken, error) {
	var refreshToken DatabaseRefreshToken

	if err := db.QueryRowContext(ctx, createRefreshTokenQuery, params.Token, params.UserID, params.ExpiresAt).Scan(
		&refreshToken.Token,
		&refreshToken.CreatedAt,
		&refreshToken.UpdatedAt,
		&refreshToken.UserID,
		&refreshToken.ExpiresAt,
		&refreshToken.RevokedAt,
	); err != nil {
		return DatabaseRefreshToken{}, fmt.Errorf("create refresh token: %w", err)
	}

	return refreshToken, nil
}

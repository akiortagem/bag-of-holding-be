package queries

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DatabaseRefreshTokenGetParams struct {
	Token string
}

type DatabaseRefreshTokenRow struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    int64
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

const getRefreshTokenQuery = `
SELECT token, created_at, updated_at, user_id, expires_at, revoked_at
FROM refresh_tokens
WHERE token = $1;
`

func GetRefreshTokenQuery(ctx context.Context, db *sql.DB, params DatabaseRefreshTokenGetParams) (DatabaseRefreshTokenRow, error) {
	var refreshToken DatabaseRefreshTokenRow

	if err := db.QueryRowContext(ctx, getRefreshTokenQuery, params.Token).Scan(
		&refreshToken.Token,
		&refreshToken.CreatedAt,
		&refreshToken.UpdatedAt,
		&refreshToken.UserID,
		&refreshToken.ExpiresAt,
		&refreshToken.RevokedAt,
	); err != nil {
		return DatabaseRefreshTokenRow{}, fmt.Errorf("get refresh token: %w", err)
	}

	return refreshToken, nil
}

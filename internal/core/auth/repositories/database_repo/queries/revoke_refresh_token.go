package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DatabaseRefreshTokenRevokeParams struct {
	Token string
}

const revokeRefreshTokenQuery = `
UPDATE refresh_tokens
SET revoked_at = NOW(),
    updated_at = NOW()
WHERE token = $1
  AND revoked_at IS NULL;
`

func RevokeRefreshTokenQuery(ctx context.Context, db *sql.DB, params DatabaseRefreshTokenRevokeParams) error {
	res, err := db.ExecContext(ctx, revokeRefreshTokenQuery, params.Token)
	if err != nil {
		return fmt.Errorf("revoke refresh token: %w", err)
	}

	if rows, err := res.RowsAffected(); err == nil && rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

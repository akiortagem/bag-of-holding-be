package queries

import (
	"context"
	"database/sql"
	"fmt"
)

type DatabaseUserGetByEmailParams struct {
	Email string
}

const getUserByEmailQuery = `
SELECT id, email, password_hash, is_active, email_verified_at, last_login_at, created_at, updated_at
FROM users
WHERE email = $1;
`

func GetUserByEmailQuery(ctx context.Context, db *sql.DB, params DatabaseUserGetByEmailParams) (DatabaseUser, error) {
	var user DatabaseUser

	if err := db.QueryRowContext(ctx, getUserByEmailQuery, params.Email).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.IsActive,
		&user.EmailVerifiedAt,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return DatabaseUser{}, fmt.Errorf("get user by email: %w", err)
	}

	return user, nil
}

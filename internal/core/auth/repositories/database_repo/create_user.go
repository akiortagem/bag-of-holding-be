package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type DatabaseUser struct {
	ID              int64
	Email           string
	PasswordHash    string
	IsActive        bool
	EmailVerifiedAt sql.NullTime
	LastLoginAt     sql.NullTime
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DatabaseUserCreateParams struct {
	Email        string
	PasswordHash string
}

const createUserQuery = `
INSERT INTO users (email, password_hash)
VALUES ($1, $2)
RETURNING id, email, password_hash, is_active, email_verified_at, last_login_at, created_at, updated_at;
`

func CreateUserQuery(ctx context.Context, db *sql.DB, params DatabaseUserCreateParams) (DatabaseUser, error) {
	var user DatabaseUser
	if err := db.QueryRowContext(ctx, createUserQuery, params.Email, params.PasswordHash).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.IsActive,
		&user.EmailVerifiedAt,
		&user.LastLoginAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return DatabaseUser{}, fmt.Errorf("create user: %w", err)
	}

	return user, nil
}

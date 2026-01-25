package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/repositories/database_repo/queries"
	"github.com/lib/pq"
)

type DBAuthService struct {
	Db *sql.DB
}

func (s *DBAuthService) CreateUser(c context.Context, newUser domain.UserCreateParams) (domain.UserCreateResponse, error) {
	dbUser, err := queries.CreateUserQuery(c, s.Db, queries.DatabaseUserCreateParams{
		Email:        newUser.Email,
		PasswordHash: newUser.Password,
	})

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return domain.UserCreateResponse{}, domain.ErrUserAlreadyExists
		}
		return domain.UserCreateResponse{}, err
	}

	resp := domain.UserCreateResponse{
		Email: dbUser.Email,
		ID:    dbUser.ID,
	}

	return resp, nil
}

func (s *DBAuthService) CreateRefreshToken(ctx context.Context, params domain.RefreshTokenCreateParams) (domain.RefreshToken, error) {
	dbToken, err := queries.CreateRefreshTokenQuery(ctx, s.Db, queries.DatabaseRefreshTokenCreateParams{
		Token:     params.Token,
		UserID:    params.UserID,
		ExpiresAt: params.ExpiresAt,
	})
	if err != nil {
		return domain.RefreshToken{}, err
	}

	var revokedAt *time.Time
	if dbToken.RevokedAt.Valid {
		revokedAt = &dbToken.RevokedAt.Time
	}

	return domain.RefreshToken{
		Token:     dbToken.Token,
		CreatedAt: dbToken.CreatedAt,
		UpdatedAt: dbToken.UpdatedAt,
		UserID:    dbToken.UserID,
		ExpiresAt: dbToken.ExpiresAt,
		RevokedAt: revokedAt,
	}, nil
}

func (s *DBAuthService) GetUserByEmail(ctx context.Context, params domain.UserGetByEmailParams) (domain.User, error) {
	dbUser, err := queries.GetUserByEmailQuery(ctx, s.Db, queries.DatabaseUserGetByEmailParams{
		Email: params.Email,
	})
	if err != nil {
		return domain.User{}, err
	}

	var emailVerifiedAt *time.Time
	if dbUser.EmailVerifiedAt.Valid {
		emailVerifiedAt = &dbUser.EmailVerifiedAt.Time
	}

	var lastLoginAt *time.Time
	if dbUser.LastLoginAt.Valid {
		lastLoginAt = &dbUser.LastLoginAt.Time
	}

	return domain.User{
		ID:              dbUser.ID,
		Email:           dbUser.Email,
		PasswordHash:    dbUser.PasswordHash,
		IsActive:        dbUser.IsActive,
		EmailVerifiedAt: emailVerifiedAt,
		LastLoginAt:     lastLoginAt,
		CreatedAt:       dbUser.CreatedAt,
		UpdatedAt:       dbUser.UpdatedAt,
	}, nil
}

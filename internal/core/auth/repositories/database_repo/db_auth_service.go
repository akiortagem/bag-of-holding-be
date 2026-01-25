package repositories

import (
	"context"
	"database/sql"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
)

type DBAuthService struct {
	Db *sql.DB
}

func (s *DBAuthService) CreateUser(c context.Context, newUser domain.UserCreateParams) (domain.UserCreateResponse, error) {
	dbUser, err := CreateUserQuery(c, s.Db, DatabaseUserCreateParams{
		Email:        newUser.Email,
		PasswordHash: newUser.Password,
	})

	if err != nil {
		return domain.UserCreateResponse{}, err
	}

	resp := domain.UserCreateResponse{
		Email: dbUser.Email,
		ID:    dbUser.ID,
	}

	return resp, nil
}

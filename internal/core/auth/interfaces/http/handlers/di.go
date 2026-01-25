package handlers

import (
	"database/sql"

	repositories "github.com/akiortagem/bag-of-holding-be/internal/core/auth/repositories/database_repo"
	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/usecase"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	"github.com/gin-gonic/gin"
)

func GetDBCreateUserHandler(db *sql.DB) func(c *gin.Context) {
	uc := usecase.CreateUserUsecase{
		Service: &repositories.DBAuthService{
			Db: db,
		},
	}

	return func(c *gin.Context) {
		CreateUserHandler(c, &uc)
	}
}

func GetDBLoginUserHandler(db *sql.DB, cfg config.ApiConfig) func(c *gin.Context) {
	uc := usecase.LoginUserUsecase{
		Service: &repositories.DBAuthService{
			Db: db,
		},
	}

	return func(c *gin.Context) {
		LoginHandler(c, &uc, cfg)
	}
}

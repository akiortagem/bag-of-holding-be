package http

import (
	"database/sql"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/interfaces/http/handlers"
	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/interfaces/http/middelwares"
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
		handlers.CreateUserHandler(c, &uc)
	}
}

func GetDBLoginUserHandler(db *sql.DB, cfg config.ApiConfig) func(c *gin.Context) {
	uc := usecase.LoginUserUsecase{
		Service: &repositories.DBAuthService{
			Db: db,
		},
	}

	return func(c *gin.Context) {
		handlers.LoginHandler(c, &uc, cfg)
	}
}

func GetJWTAuthRequiredMiddleware(cfg config.ApiConfig) gin.HandlerFunc {
	return middelwares.AuthRequired(cfg)
}

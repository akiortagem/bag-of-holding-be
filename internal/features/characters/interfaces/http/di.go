package http

import (
	"database/sql"

	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/interfaces/http/handlers"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/repositories/db"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/usecases"
	"github.com/gin-gonic/gin"
)

func GetDBCreateCharacterHandler(database *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		handlers.CreateCharacterHandler(c, &usecases.CreateCharacterUsecase{
			Service: &db.DBCharacterService{
				Db: database,
			},
		})
	}
}

func GetDBListCharactersHandler(database *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		handlers.ListCharacterHandler(c, &usecases.ListCharactersUsecase{
			Service: &db.DBCharacterService{
				Db: database,
			},
		})
	}
}

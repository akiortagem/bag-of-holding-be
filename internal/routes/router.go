package routes

import (
	"github.com/gin-gonic/gin"

	"bag-of-holding-be/internal/handlers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	healthHandler := handlers.NewHealthHandler()
	router.GET("/health", healthHandler.GetHealth)

	return router
}

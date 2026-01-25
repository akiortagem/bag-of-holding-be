package handlers

import (
	"github.com/akiortagem/bag-of-holding-be/internal/features/health/usecase"
	"github.com/gin-gonic/gin"
)

func GetHealthCheckHandler() func(c *gin.Context) {
	healthChecker := usecase.BoolHealthChecker{}
	return func(c *gin.Context) {
		CheckHealthHandler(c, &healthChecker)
	}
}

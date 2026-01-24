package handlers

import (
	"log"
	"net/http"

	"github.com/akiortagem/bag-of-holding-be/internal/features/health/di"
	"github.com/gin-gonic/gin"
)

func CheckHealthHandler(c *gin.Context, di di.HealthDI) {

	isOK := di.HealthChecker.CheckHealth()

	if !isOK {
		log.Printf("Health check is not OK")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Health check is not OK",
		})
		return
	}

	c.Status(200)
}

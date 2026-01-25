package main

import (
	"log"

	health "github.com/akiortagem/bag-of-holding-be/internal/features/health/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", health.GetHealthCheckHandler())

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

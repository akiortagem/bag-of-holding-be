package main

import (
	"log"

	"github.com/akiortagem/bag-of-holding-be/internal/core"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg := core.SetupServer()

	r.GET("/health", cfg.CheckHealth)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

package main

import (
	"database/sql"
	"log"
	"os"

	auth "github.com/akiortagem/bag-of-holding-be/internal/core/auth/interfaces/http/handlers"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	health "github.com/akiortagem/bag-of-holding-be/internal/features/health/interfaces/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	log.Printf("Setting up server")
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load .env: %v", err)
	}
	dsn := os.Getenv("DATABASE_URL")
	log.Printf("dsn is : %v", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	cfg := config.ApiConfig{
		Secret:    os.Getenv("SECRET"),
		JWTIssuer: os.Getenv("bag-of-holding"),
	}

	r := gin.Default()

	r.GET("/health", health.GetHealthCheckHandler())
	r.POST("/api/users", auth.GetDBCreateUserHandler(db))
	r.POST("/api/login", auth.GetDBLoginUserHandler(db, cfg))

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

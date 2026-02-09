package main

import (
	"database/sql"
	"log"
	"os"

	auth "github.com/akiortagem/bag-of-holding-be/internal/core/auth/interfaces/http"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	characters "github.com/akiortagem/bag-of-holding-be/internal/features/characters/interfaces/http"
	health "github.com/akiortagem/bag-of-holding-be/internal/features/health/interfaces/http/handlers"
	"github.com/gin-contrib/cors"
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
		JWTIssuer: "bag-of-holding",
		Domain:    os.Getenv("DOMAIN"),
		Platform:  os.Getenv("PLATFORM"),
	}

	r := gin.Default()
	// r.Use(cors.Default())
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173", "http://0.0.0.0:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	r.GET("/health", health.GetHealthCheckHandler())
	r.POST("/api/users", auth.GetDBCreateUserHandler(db))
	r.POST("/api/login", auth.GetDBLoginUserHandler(db, cfg))
	r.POST("/api/refresh", auth.GetDBRefreshTokenHandler(db, cfg))
	r.POST("/api/characters", auth.GetJWTAuthRequiredMiddleware(cfg), characters.GetDBCreateCharacterHandler(db))
	r.GET("/api/characters", auth.GetJWTAuthRequiredMiddleware(cfg), characters.GetDBListCharactersHandler(db))
	r.GET("/protected-health", auth.GetJWTAuthRequiredMiddleware(cfg), health.GetHealthCheckHandler())

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

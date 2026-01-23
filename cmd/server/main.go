package main

import (
	"log"

	"bag-of-holding-be/internal/routes"
)

func main() {
	router := routes.NewRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

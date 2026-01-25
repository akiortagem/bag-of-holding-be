package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/usecase"
	"github.com/gin-gonic/gin"
)

func CreateUserHandler(c *gin.Context, creator *usecase.CreateUserUsecase) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	resp, err := creator.CreateUser(c.Request.Context(), payload)
	if err != nil {
		if strings.Contains(err.Error(), "decode create user params") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user already exists"})
			return
		}
		log.Printf("Failed creating user: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

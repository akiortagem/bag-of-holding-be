package handlers

import (
	"errors"
	"log"
	"net/http"
	"strings"

	authDom "github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/usecases"
	"github.com/gin-gonic/gin"
)

func CreateCharacterHandler(c *gin.Context, creator *usecases.CreateCharacterUsecase) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	val, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := val.(int64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	resp, err := creator.CreateCharacter(c.Request.Context(), payload, userID)
	if err != nil {
		if strings.Contains(err.Error(), "decode create character params") {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
			return
		}

		if errors.Is(err, authDom.ErrUserNotAuthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		log.Printf("Failed creating user: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

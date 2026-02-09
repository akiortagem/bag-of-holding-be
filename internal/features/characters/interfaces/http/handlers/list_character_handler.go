package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	authDom "github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	coreDom "github.com/akiortagem/bag-of-holding-be/internal/core/domain"
	charDom "github.com/akiortagem/bag-of-holding-be/internal/features/characters/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/features/characters/usecases"
	"github.com/gin-gonic/gin"
)

func ListCharacterHandler(c *gin.Context, lister *usecases.ListCharactersUsecase) {
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

	page, err := parseQueryInt(c, "page", 1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page"})
		return
	}

	pageSize, err := parseQueryInt(c, "page_size", 20)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page_size"})
		return
	}

	keyword := c.Query("keyword")
	params := charDom.ListCharacterParams{
		UserID:   userID,
		Page:     page,
		PageSize: pageSize,
	}
	if keyword != "" {
		params.Keyword = &keyword
	}

	payload, err := json.Marshal(params)
	if err != nil {
		log.Printf("Failed list characters payload: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	resp, err := lister.ListCharacters(c.Request.Context(), payload, userID)
	if err != nil {
		if errors.Is(err, authDom.ErrUserNotAuthorized) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
			return
		}

		log.Printf("Failed listing characters: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list characters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resp,
		"page": coreDom.PaginatedData{
			Page:     page,
			PageSize: pageSize,
		},
	})
}

func parseQueryInt(c *gin.Context, key string, def int) (int, error) {
	value := c.Query(key)
	if value == "" {
		return def, nil
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	return parsed, nil
}

package handlers

import (
	"errors"
	"net/http"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/domain"
	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/usecase"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context, loginer *usecase.LoginUserUsecase, cfg config.ApiConfig) {
	payload, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	var resp domain.UserLoginResponse

	if resp, err = loginer.LoginUser(c, payload, cfg); err != nil {
		if errors.Is(err, &usecase.LoginFailedError{}) {
			c.Status(401)
			return
		}
		c.Status(500)
		return
	}

	c.JSON(200, resp)
}

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

	needHttps := false

	if !cfg.IsDebug() {
		needHttps = true
	}

	c.SetSameSite(http.SameSiteLaxMode)

	c.SetCookie(
		"refresh",
		resp.RefreshToken,
		3600*24*7, // one week
		"/api/refresh",
		"",
		needHttps,
		true,
	)

	c.JSON(200, resp)
}

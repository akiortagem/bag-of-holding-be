package handlers

import (
	"errors"
	"net/http"

	"github.com/akiortagem/bag-of-holding-be/internal/core/auth/usecase"
	"github.com/akiortagem/bag-of-holding-be/internal/core/config"
	"github.com/gin-gonic/gin"
)

func RefreshTokenHandler(c *gin.Context, refresher *usecase.RefreshTokenUsecase, cfg config.ApiConfig) {
	refreshToken, err := c.Cookie("refresh")
	if err != nil || refreshToken == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	resp, err := refresher.RefreshToken(c.Request.Context(), refreshToken, cfg)
	if err != nil {
		var invalidErr *usecase.RefreshTokenInvalidError
		if errors.As(err, &invalidErr) {
			c.Status(http.StatusUnauthorized)
			return
		}
		c.Status(http.StatusInternalServerError)
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

	c.JSON(http.StatusOK, resp)
}

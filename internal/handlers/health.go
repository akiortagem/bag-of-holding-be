package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"bag-of-holding-be/internal/usecase"
)

type HealthHandler struct {
	usecase usecase.HealthUsecase
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{usecase: usecase.NewHealthUsecase()}
}

func (handler *HealthHandler) GetHealth(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"status": handler.usecase.Status()})
}

package core

import (
	hdi "github.com/akiortagem/bag-of-holding-be/internal/features/health/di"
	hh "github.com/akiortagem/bag-of-holding-be/internal/features/health/interfaces/http/handlers"
	hu "github.com/akiortagem/bag-of-holding-be/internal/features/health/usecase"
	"github.com/gin-gonic/gin"
)

type DefaultHandlerCfg struct {
	CheckHealth func(c *gin.Context)
}

func SetupServer() DefaultHandlerCfg {
	healthDI := hdi.HealthDI{
		HealthChecker: &hu.BoolHealthChecker{},
	}

	cfg := DefaultHandlerCfg{
		CheckHealth: func(c *gin.Context) {
			hh.CheckHealthHandler(c, healthDI)
		},
	}

	return cfg
}

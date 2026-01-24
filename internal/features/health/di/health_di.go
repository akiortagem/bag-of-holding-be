package di

import "github.com/akiortagem/bag-of-holding-be/internal/features/health/usecase"

type HealthDI struct {
	HealthChecker usecase.HealthChecker
}

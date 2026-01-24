package usecase

type HealthChecker interface {
	CheckHealth() bool
}

type BoolHealthChecker struct{}

func (u *BoolHealthChecker) CheckHealth() bool {
	return true
}

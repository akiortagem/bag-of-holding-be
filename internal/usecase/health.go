package usecase

type HealthUsecase interface {
	Status() string
}

type healthUsecase struct{}

func NewHealthUsecase() HealthUsecase {
	return &healthUsecase{}
}

func (usecase *healthUsecase) Status() string {
	return "ok"
}

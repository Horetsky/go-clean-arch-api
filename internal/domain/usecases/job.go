package usecases

type JobUsecase interface {
}

type jobUsecase struct {
}

func NewJobUsecase() JobUsecase {
	return &jobUsecase{}
}

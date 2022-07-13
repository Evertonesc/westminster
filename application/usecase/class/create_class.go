package usecase

import "context"

type CreateClassUsecase struct {
}

func New() CreateClassUsecase {
	return CreateClassUsecase{}
}

func (uc CreateClassUsecase) CreateClass(ctx context.Context, className string) error {
	return nil
}

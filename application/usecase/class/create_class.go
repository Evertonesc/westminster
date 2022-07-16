package usecase

import (
	"context"
	"orinz/application/domain/class"
)

type CreateClassUsecase struct {
	r class.Repository
}

func New(r class.Repository) CreateClassUsecase {
	return CreateClassUsecase{
		r: r,
	}
}

func (uc CreateClassUsecase) Execute(ctx context.Context, className class.Class) error {
	return uc.r.Create(ctx, className)
}

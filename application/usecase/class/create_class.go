package usecase

import (
	"context"
	"orinz/application/domain/class"
)

type CreateClassUseCase struct {
	r class.Repository
}

func NewCreateClass(r class.Repository) CreateClassUseCase {
	return CreateClassUseCase{
		r: r,
	}
}

func (uc CreateClassUseCase) Execute(ctx context.Context, className class.Class) error {
	return uc.r.Create(ctx, className)
}

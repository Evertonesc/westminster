package usecase

import "context"

type CreateMemberUseCase struct {
}

func NewCreateMemberUseCase() CreateMemberUseCase {
	return CreateMemberUseCase{}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, member any) error {
	return nil
}

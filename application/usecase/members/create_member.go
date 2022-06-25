package usecase

import (
	"context"
	"orinz/application/adapter/rest/presenter"
)

type CreateMemberUseCase struct {
}

func NewCreateMemberUseCase() CreateMemberUseCase {
	return CreateMemberUseCase{}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, member presenter.MemberRequest) error {
	return nil
}

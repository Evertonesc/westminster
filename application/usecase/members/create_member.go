package usecase

import (
	"context"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
)

type CreateMemberUseCase struct {
	repository member.Repository
}

func NewCreateMemberUseCase(repository member.Repository) CreateMemberUseCase {
	return CreateMemberUseCase{
		repository: repository,
	}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, memberRequest presenter.MemberRequest) error {

	return nil
}

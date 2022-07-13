package usecase

import (
	"context"
	"errors"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
)

type CreateMemberUseCase struct {
	repository member.Repository
}

func New(repository member.Repository) CreateMemberUseCase {
	return CreateMemberUseCase{
		repository: repository,
	}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, memberRequest presenter.MemberRequest) error {

	member := member.New(memberRequest.Name, memberRequest.BirthDate, memberRequest.Address, memberRequest.Email)
	if !member.Validate() {
		return errors.New("member data is invalid")
	}

	if err := uc.repository.CreateMember(ctx, member); err != nil {
		return err
	}

	return nil
}

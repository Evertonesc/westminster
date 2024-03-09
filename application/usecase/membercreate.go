package usecase

import (
	"context"

	"westminster/application/domain"
)

type CreateMemberUseCase struct {
	w domain.MemberWriter
}

func NewCreateMemberUseCase(w domain.MemberWriter) *CreateMemberUseCase {
	return &CreateMemberUseCase{
		w: w,
	}
}

func (uc *CreateMemberUseCase) Execute(
	ctx context.Context,
	i MemberInteractor,
) (*domain.Member, error) {
	member := i.ToDomain()

	return uc.w.CreateMember(ctx, member)
}

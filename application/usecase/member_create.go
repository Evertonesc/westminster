package usecase

import (
	"context"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
)

type CreateMemberUseCase struct {
	w member.Writer
}

func NewCreateMemberUseCase(w member.Writer) CreateMemberUseCase {
	return CreateMemberUseCase{
		w: w,
	}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, mr presenter.MemberRequest) (member.Member, error) {
	var m = member.NewMember(mr.Name, mr.Location, mr.FinancialNumber, mr.Enabled)
	return uc.w.CreateMember(ctx, m)
}

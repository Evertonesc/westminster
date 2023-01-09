package usecase

import (
	"context"
	"westminster/application/adapter/rest/presenter"
	"westminster/application/domain"
)

type CreateMemberUseCase struct {
	w domain.MemberWriter
}

func NewCreateMemberUseCase(w domain.MemberWriter) CreateMemberUseCase {
	return CreateMemberUseCase{
		w: w,
	}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, mr presenter.MemberRequest) (domain.Member, error) {
	var m = domain.NewMember(mr.Name, mr.Location, mr.FinancialNumber, mr.Enabled)
	return uc.w.CreateMember(ctx, m)
}

package usecase

import (
	"context"
	"westminster/application/domain"
)

type CreateMemberUseCase struct {
	w domain.MemberWriter
}

type MemberInteractor struct {
	Name            string
	Location        string
	FinancialNumber int
	Enabled         bool
}

func NewCreateMemberUseCase(w domain.MemberWriter) CreateMemberUseCase {
	return CreateMemberUseCase{
		w: w,
	}
}

func (uc CreateMemberUseCase) Execute(ctx context.Context, i MemberInteractor) (domain.Member, error) {
	var m = domain.NewMember(i.Name, i.Location, i.FinancialNumber, i.Enabled)
	return uc.w.CreateMember(ctx, m)
}

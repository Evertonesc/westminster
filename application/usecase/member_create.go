package usecase

import (
	"context"
	"time"
	"westminster/application/domain"
)

type CreateMemberUseCase struct {
	w domain.MemberWriter
}

type MemberInteractor struct {
	ID              string
	Name            string
	Location        string
	FinancialNumber int
	Enabled         bool
	CreateAt        time.Time
	UpdatedAt       time.Time
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

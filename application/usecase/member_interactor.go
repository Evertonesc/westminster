package usecase

import (
	"time"

	"westminster/application/domain"
)

type MemberInteractor struct {
	ID              string
	Name            string
	Location        string
	FinancialNumber int
	Enabled         bool
	CreateAt        time.Time
	UpdatedAt       time.Time
}

func (mi *MemberInteractor) ToDomain() *domain.Member {
	return domain.NewMember(mi.Name, mi.Location, mi.FinancialNumber, mi.Enabled)
}

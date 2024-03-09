package domain

import (
	"context"
	"time"
)

type MemberWriter interface {
	CreateMember(ctx context.Context, member *Member) (*Member, error)
}

type Member struct {
	ID              string
	Name            string
	FinancialNumber int
	Location        string
	Enabled         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewMember(name, location string, financialNumber int, enabled bool) *Member {
	return &Member{
		Name:            name,
		FinancialNumber: financialNumber,
		Location:        location,
		Enabled:         enabled,
	}
}

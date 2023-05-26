package handler

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"context"

	"westminster/application/domain"
	"westminster/application/usecase"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, i usecase.MemberInteractor) (domain.Member, error)
}

type MemberHandler struct {
	uc CreateMemberUseCase
}

func NewMemberHandler(uc CreateMemberUseCase) MemberHandler {
	return MemberHandler{
		uc: uc,
	}
}

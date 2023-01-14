package handler

import (
	"context"
	"westminster/application/domain"
	"westminster/application/usecase"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, i usecase.MemberInteractor) (domain.Member, error)
}

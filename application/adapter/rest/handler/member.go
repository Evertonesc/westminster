package handler

import (
	"context"
	"westminster/application/adapter/rest/presenter"
	"westminster/application/domain"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, mr presenter.MemberRequest) (domain.Member, error)
}

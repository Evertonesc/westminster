package handler

import (
	"context"
	"westminster/application/adapter/rest/presenter"
	"westminster/application/domain/member"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, mr presenter.MemberRequest) (member.Member, error)
}

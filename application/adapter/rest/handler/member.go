package handler

import (
	"context"
	"orinz/application/adapter/rest/presenter"
	"orinz/application/domain/member"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, mr presenter.MemberRequest) (member.Member, error)
}

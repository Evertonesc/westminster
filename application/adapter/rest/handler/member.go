package handler

import (
	"context"
	"orinz/application/adapter/rest/presenter"
)

type CreateMemberUseCase interface {
	Execute(ctx context.Context, mr presenter.MemberRequest) error
}

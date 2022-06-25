package usecase

import (
	"context"
	"orinz/application/adapter/rest/presenter"
)

type CreateMember interface {
	Execute(ctx context.Context, memberRequest presenter.MemberRequest) error
}

package usecase

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"context"
	"orinz/application/adapter/rest/presenter"
)

type CreateMember interface {
	Execute(ctx context.Context, memberRequest presenter.MemberRequest) error
}

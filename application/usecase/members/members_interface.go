package usecase

import "context"

type CreateMember interface {
	Execute(ctx context.Context, member any) error
}

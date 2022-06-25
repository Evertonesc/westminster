package member

import "context"

type Repository interface {
	CreateMember(ctx context.Context, member Member) error
}

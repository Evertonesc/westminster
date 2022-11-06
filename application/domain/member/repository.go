package member

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import "context"

type Repository interface {
	CreateMember(ctx context.Context, member Member) (Member, error)
}

package class

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import "context"

type Repository interface {
	Create(ctx context.Context, class Class) error
	Get(ctx context.Context, ID string) error
}

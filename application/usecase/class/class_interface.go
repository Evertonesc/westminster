package usecase

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import "context"

type CreateClass interface {
	CreateClass(ctx context.Context, className string) error
}

package usecase

//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE

import (
	"context"
	"orinz/application/domain/class"
)

type CreateClass interface {
	Execute(ctx context.Context, className class.Class) error
}

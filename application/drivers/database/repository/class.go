package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"orinz/application/domain/class"
)

type Class struct {
	ID   string
	Name string
}

func NewClass(class class.Class) Class {
	return Class{
		ID:   uuid.NewString(),
		Name: class.Name,
	}
}

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return ClassRepository{
		db: db,
	}
}

func (r ClassRepository) Create(ctx context.Context, class class.Class) error {
	return nil
}

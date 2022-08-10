package repository

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"orinz/application/domain/class"
	"time"
)

type Class struct {
	ID        string
	Name      string
	CreateAt  time.Time
	UpdatedAt time.Time
}

func NewClass(class class.Class) Class {
	return Class{
		ID:        uuid.NewString(),
		Name:      class.Name,
		CreateAt:  time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
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
	c := NewClass(class)
	return r.db.Create(c).Error
}

func (r ClassRepository) Get(ctx context.Context, ID string) error {
	return nil
}

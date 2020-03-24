package model

import (
	"context"
	"time"

	"github.com/go-pg/pg/orm"
)

// Model contains common fields for all tables
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Model) BeforeInsert(_ context.Context, _ orm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

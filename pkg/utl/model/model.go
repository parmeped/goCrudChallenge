package model

import (
	"context"
	"time"

	"github.com/go-pg/pg/orm"
)

// Base contains common fields for all tables
type Base struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(_ context.Context, _ orm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(_ context.Context, _ orm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}

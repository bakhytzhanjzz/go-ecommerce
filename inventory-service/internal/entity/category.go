package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name        string         `gorm:"size:255;not null;unique" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

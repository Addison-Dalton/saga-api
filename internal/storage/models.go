package storage

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // hide DeletedAt from JSON responses
}

type Character struct {
	BaseModel

	Name string `gorm:"uniqueIndex;not null" json:"name"`
	HP   int    `gorm:"default:100" json:"hp"`
	Mana int    `gorm:"default:100" json:"mana"`
}

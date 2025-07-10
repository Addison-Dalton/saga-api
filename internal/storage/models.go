package storage

import "gorm.io/gorm"

type Character struct {
	gorm.Model

	Name string `gorm:"uniqueIndex;not null" json:"name"`
	HP   int    `gorm:"default:100" json:"hp"`
	Mana int    `gorm:"default:100" json:"mana"`
}

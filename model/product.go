package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}

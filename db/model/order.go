package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint `gorm:"not null;index"` // Foreign key
	Description string
	OrderDate   string
	User        User `gorm:"foreignKey:UserID"` // Establishes the foreign key relationship
}

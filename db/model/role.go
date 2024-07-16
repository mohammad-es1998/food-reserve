package model

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Permissions string // Comma-separated list of permissions
}

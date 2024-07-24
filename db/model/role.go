package model

import (
	"gorm.io/gorm"
	"strings"
)

type Role struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Permissions string // Comma-separated list of permissions
}

// Method to check if the role has a specific permission
func (r *Role) HasPermission(permission string) bool {
	permissions := strings.Split(r.Permissions, ",")
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}

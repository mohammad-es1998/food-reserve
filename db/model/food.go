package model

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	name  string
	price string
}

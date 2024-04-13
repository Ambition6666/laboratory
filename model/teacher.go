package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	User
}
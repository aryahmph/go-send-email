package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"notNull,unique"`
}

package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"primaryKey";json:"username"`
	Password string `gorm:"unique;not null";json:"password"`
}

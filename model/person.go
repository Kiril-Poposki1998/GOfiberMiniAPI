package model

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	id      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

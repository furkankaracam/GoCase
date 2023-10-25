package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	Name         string
	Surname      string
	SchoolNumber string
}

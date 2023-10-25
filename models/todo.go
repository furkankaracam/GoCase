package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Id        uint `gorm:"primaryKey"`
	StudentID uint `gorm:"foreignKey:ID"`
	Title     string
	Date      string
	StartTime string
	EndTime   string
	Status    int
	Student   Student
}

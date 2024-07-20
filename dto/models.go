package dto

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Id          uint `gorm:"default:primaryKey"`
	Title       string
	Description string
	DueDate     string
	CreatedAt   string
	UpdatedAt   string
}

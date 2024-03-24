package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title     string `json:"title" validate:"required"`
	Completed *bool  `json:"completed" validate:"required,boolean"`
}

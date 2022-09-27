package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Author      string `json:"author" form:"author"`
	Publisher   string `json:"publisher" form:"publisher"`
}

package model

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

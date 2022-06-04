package repository

import (
	"todoapp-backend/app/model"

	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

func NewUsers(db *gorm.DB) *Users {
	return &Users{db}
}

func (u *Users) List() []model.User {
	users := []model.User{}
	u.DB.Find(&users)
	return users
}

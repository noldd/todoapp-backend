package repository

import (
	"log"
	"todoapp-backend/app/model"

	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

func NewUsers(db *gorm.DB) *Users {
	return &Users{db}
}

func (r *Users) List() []model.User {
	users := []model.User{}
	r.DB.Find(&users)
	return users
}

func (r *Users) Create(user model.User) (model.User, error) {
	err := wrapError(r.DB.Save(&user).Error)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *Users) EmailExists(email string) bool {
	users := []model.User{}
	err := r.DB.Where("email = ?", email).Limit(1).Find(&users).Error
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	if len(users) > 0 {
		return true
	}
	return false
}

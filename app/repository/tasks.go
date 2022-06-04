package repository

import (
	"todoapp-backend/app/model"

	"gorm.io/gorm"
)

type Tasks struct {
	DB *gorm.DB
}

func NewTasks(db *gorm.DB) *Tasks {
	return &Tasks{db}
}

func (r *Tasks) List() []model.Task {
	tasks := []model.Task{}
	r.DB.Find(&tasks)
	return tasks
}

func (r *Tasks) GetById(id uint) (model.Task, error) {
	task := model.Task{}

	err := wrapError(r.DB.First(&task, id).Error)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *Tasks) Create(task model.Task) (model.Task, error) {
	err := wrapError(r.DB.Save(&task).Error)
	if err != nil {
		return task, err
	}

	return task, nil
}

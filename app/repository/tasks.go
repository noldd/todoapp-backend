package repository

import (
	"todoapp-backend/app/model"

	"gorm.io/gorm"
)

// TODO: Error handling for all methods
type Tasks struct {
	DB *gorm.DB
}

func NewTasks(db *gorm.DB) *Tasks {
	return &Tasks{db}
}

func (t *Tasks) List() []model.Task {
	tasks := []model.Task{}
	t.DB.Find(&tasks)
	return tasks
}

func (t *Tasks) GetById(id uint) (model.Task, error) {
	task := model.Task{}

	err := t.DB.First(&task, id).Error
	if err != nil {
		return task, err
	}

	return task, nil
}

func (t *Tasks) Create(task model.Task) (model.Task, error) {
	err := t.DB.Save(&task).Error

	if err != nil {
		return task, err
	}

	return task, nil
}

package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"todoapp-backend/app/model"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Tasks struct {
	DB *gorm.DB
}

func (t *Tasks) Router(r chi.Router) {
	r.Get("/", t.List)
	r.Post("/", t.Post)
}

func NewTasksRouter(db *gorm.DB) *Tasks {
	return &Tasks{db}
}

func (t *Tasks) List(w http.ResponseWriter, r *http.Request) {
	tasks := []model.Task{}
	t.DB.Find(&tasks)
	respondJSON(w, http.StatusOK, tasks)
}

func (t *Tasks) Post(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		// TODO: Respond with error
		log.Fatal("Failed to decode json")
		return
	}
	defer r.Body.Close()

	if err := t.DB.Save(&task).Error; err != nil {
		// TODO: Respond with error
		return
	}
	respondJSON(w, http.StatusCreated, task)
}
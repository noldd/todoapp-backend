package routes

import (
	"errors"
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
	r.Get("/{id}", t.Get)
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

func (t *Tasks) Get(w http.ResponseWriter, r *http.Request) {
	task := model.Task{}
	if err := t.DB.First(&task, chi.URLParam(r, "id")).Error; err != nil {
		switch {
		case errors.Is(err, gorm.ErrRecordNotFound):
			respondError(w, http.StatusNotFound, "Task not found")
			return
		default:
			respondError(w, http.StatusInternalServerError, "Internal server error")
			log.Fatalf("Error while getting task %v: %c", task.ID, err)
			return
		}
	}

	respondJSON(w, http.StatusOK, task)
}

func (t *Tasks) Post(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := parseJSON(r.Body, &task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	if err := t.DB.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, "Internal server error")
		log.Printf("Failed to save post to DB: %s", err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, task)
}

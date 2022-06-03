package controller

import (
	"errors"
	"log"
	"net/http"
	"todoapp-backend/app/model"
	"todoapp-backend/app/repository"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Tasks struct {
	Repository *repository.Tasks
}

func (t *Tasks) Router(r chi.Router) {
	r.Get("/", t.List)
	r.Get("/{id}", t.Get)
	r.Post("/", t.Post)
}

func NewTasksController(repository *repository.Tasks) *Tasks {
	return &Tasks{repository}
}

func (t *Tasks) List(w http.ResponseWriter, r *http.Request) {
	tasks := t.Repository.List()
	respondJSON(w, http.StatusOK, tasks)
}

func (t *Tasks) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		// TODO: Better response?
		respondError(w, http.StatusBadRequest, "Bad request")
	}

	task, err := t.Repository.GetById(id)
	if err != nil {
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

	task, err := t.Repository.Create(task)

	if err != nil {
		respondError(w, http.StatusInternalServerError, "Internal server error")
		log.Printf("Failed to save post to DB: %s", err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, task)
}

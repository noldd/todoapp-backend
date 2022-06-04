package controller

import (
	"net/http"
	"todoapp-backend/app/model"
	"todoapp-backend/app/repository"

	"github.com/go-chi/chi/v5"
)

type Tasks struct {
	Repository *repository.Tasks
}

func NewTasks(repository *repository.Tasks) *Tasks {
	return &Tasks{repository}
}

func (c *Tasks) Router(r chi.Router) {
	r.Get("/", c.List)
	r.Get("/{id}", c.Get)
	r.Post("/", c.Post)
}

func (c *Tasks) List(w http.ResponseWriter, r *http.Request) {
	tasks := c.Repository.List()
	respondJSON(w, http.StatusOK, tasks)
}

func (c *Tasks) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, err)
		return
	}

	task, err := c.Repository.GetById(id)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJSON(w, http.StatusOK, task)
}

func (c *Tasks) Post(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	if err := parseJSON(r.Body, &task); err != nil {
		respondError(w, err)
		return
	}

	task, err := c.Repository.Create(task)
	if err != nil {
		respondError(w, err)
		return
	}

	respondJSON(w, http.StatusCreated, task)
}

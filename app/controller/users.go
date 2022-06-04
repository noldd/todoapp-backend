package controller

import (
	"net/http"
	"todoapp-backend/app/repository"

	"github.com/go-chi/chi/v5"
)

type Users struct {
	Repository *repository.Users
}

func NewUsers(repository *repository.Users) *Users {
	return &Users{repository}
}

func (c *Users) Router(r chi.Router) {
	r.Get("/", c.List)
}

func (c *Users) List(w http.ResponseWriter, r *http.Request) {
	users := c.Repository.List()
	respondJSON(w, http.StatusOK, users)
}

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

func (u *Users) Router(r chi.Router) {
	r.Get("/", u.List)
}

func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	users := u.Repository.List()
	respondJSON(w, http.StatusOK, users)
}

package controller

import (
	"net/http"
	"todoapp-backend/app/model"
	"todoapp-backend/app/repository"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Repository *repository.Users
}

func NewUsers(repository *repository.Users) *Users {
	return &Users{repository}
}

func (c *Users) Router(r chi.Router) {
	r.Get("/", c.List)
	r.Post("/", c.Register)
}

func (c *Users) List(w http.ResponseWriter, r *http.Request) {
	users := c.Repository.List()
	respondJSON(w, http.StatusOK, users)
}

func (c *Users) Login(w http.ResponseWriter, r *http.Request) {}

func (c *Users) Register(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := parseJSON(r.Body, &user); err != nil {
		respondError(w, err)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		respondError(w, err)
		return
	}
	user.Password = string(password)

	// Make sure the user doesn't exist
	// TODO: It's faster to check the database error than to check wheter the
	// email exists or not. We could save an entire query here
	if c.Repository.EmailExists(user.Email) {
		respondError(w, newErrEmailInUse())
		return
	}

	user, err = c.Repository.Create(user)
	if err != nil {
		// TODO: Check the error handling here
		respondError(w, err)
		return
	}

	// TODO: Omit password from response
	respondJSON(w, http.StatusCreated, user)
}

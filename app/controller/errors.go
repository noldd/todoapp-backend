package controller

import (
	"net/http"
)

type controllerError struct {
	status  int
	message string
}

func (e controllerError) APIError() (int, string) {
	return e.status, e.message
}

func (e controllerError) Error() string {
	return e.message
}

func newErrEmailInUse() controllerError {
	return controllerError{http.StatusConflict, "Email already in use"}
}

func wrapError(err error) error {
	switch {
	default:
		return err
	}
}

var (
	JSONParseError = &controllerError{http.StatusBadRequest, "Failed to parse request body"}
)

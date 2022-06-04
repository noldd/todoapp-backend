package ApiErrors

// Inspired by: https://www.joeshaw.org/error-handling-in-go-http-applications/

import "net/http"

type APIError interface {
	APIError() (int, string)
}

type sentinelAPIError struct {
	status  int
	message string
}

func (e sentinelAPIError) Error() string {
	return e.message
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.message
}

type sentinelWrappedError struct {
	error
	sentinel *sentinelAPIError
}

func (e sentinelWrappedError) Is(err error) bool {
	return e.sentinel == err
}

func (e sentinelWrappedError) APIError() (int, string) {
	return e.sentinel.APIError()
}

func WrapError(err error, sentinel *sentinelAPIError) error {
	return sentinelWrappedError{err, sentinel}
}

var (
	ErrNotFound   = &sentinelAPIError{http.StatusNotFound, "Not found"}
	ErrInternal   = &sentinelAPIError{http.StatusInternalServerError, "Internal server error"}
	ErrBadRequest = &sentinelAPIError{http.StatusBadRequest, "Bad request"}
)

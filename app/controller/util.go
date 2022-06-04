package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"todoapp-backend/app/ApiErrors"
)

func parseJSON(r io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(target); err != nil {
		return JSONParseError
	}
	return nil
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, err error) {
	var apiErr ApiErrors.APIError
	if errors.As(err, &apiErr) {
		status, message := apiErr.APIError()
		respondJSON(w, status, map[string]string{"error": message})
		return
	}

	// Default to internal server error
	status, message := ApiErrors.ErrInternal.APIError()
	respondJSON(w, status, map[string]string{"error": message})
}

// TODO: Return better error?
func parseID(in string) (uint, error) {
	outU64, err := strconv.ParseUint(in, 10, 32)

	if err != nil {
		return uint(outU64), nil
	}

	return uint(outU64), nil
}

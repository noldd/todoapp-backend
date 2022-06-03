package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// TODO: Error message that can be sent to the client
func parseJSON(r io.ReadCloser, target interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(target); err != nil {
		return err
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

func respondError(w http.ResponseWriter, status int, message string) {
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

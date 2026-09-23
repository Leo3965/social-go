package api

import (
	"encoding/json"
	"net/http"
)

const MaxBytes = 1_048_578 // 1mb

type errorDto struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	r.Body = http.MaxBytesReader(w, r.Body, int64(MaxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) error {
	return writeJSON(w, status, &errorDto{
		Error: message,
	})
}

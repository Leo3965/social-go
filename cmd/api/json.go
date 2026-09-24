package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

const MaxBytes = 1_048_578 // 1mb

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type errorDto struct {
	Error string `json:"error"`
}

type applicationHttpEnvelope struct {
	Data any `json:"data"`
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

func responseJSON(w http.ResponseWriter, status int, data any) error {
	return writeJSON(w, status, &applicationHttpEnvelope{
		Data: data,
	})
}

package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, TextPlain)
	_, _ = w.Write([]byte("ok"))
}

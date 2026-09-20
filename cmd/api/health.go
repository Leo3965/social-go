package api

import (
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(ContentType, TextPlain)
	app.Store.Posts().Create(r.Context())
	_, _ = w.Write([]byte("ok"))
}

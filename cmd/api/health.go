package api

import (
	"net/http"
)

func (app *Application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status":  "ok",
		"env":     app.Config.Env,
		"version": app.Config.Version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

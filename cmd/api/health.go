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
		_ = writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

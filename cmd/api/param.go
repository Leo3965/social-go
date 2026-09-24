package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (app *Application) getIDParam(r *http.Request, param string) (int64, error) {
	idParam := chi.URLParam(r, param)
	return strconv.ParseInt(idParam, 10, 64)
}

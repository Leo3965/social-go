package api

import (
	"errors"
	"net/http"

	"github.com/Leo3965/social/internal/store"
)

func (app *Application) findUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r, "userID")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	user, err := app.Store.Users().Find(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}
		return
	}

	if err = responseJSON(w, http.StatusOK, user); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

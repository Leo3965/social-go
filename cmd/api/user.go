package api

import (
	"context"
	"errors"
	"net/http"

	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
)

type userKey string

const userCtx userKey = "user"

func (app *Application) findUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)
	if err := responseJSON(w, http.StatusOK, user); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *Application) followUserHandler(w http.ResponseWriter, r *http.Request) {
	follower := getUserFromContext(r)
	var userID int64 = 1

	ctx := r.Context()

	err := app.Store.Followers().Follow(ctx, userID, follower.ID)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrConflict):
			app.conflictResponse(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}
		return
	}

	if err = responseJSON(w, http.StatusNoContent, nil); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *Application) unfollowUserHandler(w http.ResponseWriter, r *http.Request) {
	unfollower := getUserFromContext(r)
	var userID int64 = 1

	ctx := r.Context()

	err := app.Store.Followers().Unfollow(ctx, userID, unfollower.ID)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err = responseJSON(w, http.StatusNoContent, nil); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func getUserFromContext(r *http.Request) *model.User {
	user, _ := r.Context().Value(userCtx).(*model.User)
	return user
}

func (app *Application) userContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

		ctx = context.WithValue(ctx, userCtx, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

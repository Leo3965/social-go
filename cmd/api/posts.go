package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Leo3965/social/cmd/api/dto"
	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
	"github.com/go-chi/chi/v5"
)

func (app *Application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &model.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.Store.Posts().Create(ctx, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}
}

func (app *Application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.Store.Posts().FindById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

}

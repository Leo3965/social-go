package api

import (
	"errors"
	"net/http"

	"github.com/Leo3965/social/cmd/api/dto"
	"github.com/Leo3965/social/internal/data/model"
	"github.com/Leo3965/social/internal/store"
)

func (app *Application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload dto.CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
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

func (app *Application) findByIDPostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r, "postID")
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	post, err := app.Store.Posts().Find(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}
		return
	}

	comments, err := app.Store.Comments().FindByPostId(ctx, id)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	post.Comments = comments

	if err := writeJSON(w, http.StatusOK, post); err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

}

func (app *Application) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.getIDParam(r, "postID")
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	ctx := r.Context()

	if err = app.Store.Posts().Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalErrorResponse(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (app *Application) patchPostHandler(w http.ResponseWriter, r *http.Request) {
	return
}

package api

import (
	"net/http"

	"github.com/Leo3965/social/internal/data/application"
)

func (app *Application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// pagination, search and filter
	fq := application.PaginatedFeedQuery{
		Limit:  10,
		Offset: 0,
		Sort:   "desc",
	}

	fq, err := fq.Parse(r)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err = Validate.Struct(fq); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	ctx := r.Context()

	feed, err := app.Store.Posts().GetUserFeed(ctx, int64(1), fq)
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err = responseJSON(w, http.StatusOK, feed); err != nil {
		app.internalErrorResponse(w, r, err)
	}
}

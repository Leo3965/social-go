package api

import "net/http"

func (app *Application) getUserFeedHandler(w http.ResponseWriter, r *http.Request) {
	// pagination, search and filter

	ctx := r.Context()

	feed, err := app.Store.Posts().GetUserFeed(ctx, int64(1))
	if err != nil {
		app.internalErrorResponse(w, r, err)
		return
	}

	if err = responseJSON(w, http.StatusOK, feed); err != nil {
		app.internalErrorResponse(w, r, err)
	}
}

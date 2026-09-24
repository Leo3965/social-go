package api

import (
	"log"
	"net/http"

	"github.com/Leo3965/social/internal/store"
)

func (app *Application) internalErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("internal server error: %s path: %s error: %s", r.Method, r.URL.Path, err)

	if err := writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem"); err != nil {
		log.Println(err)
	}
}

func (app *Application) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("bad request: %s path: %s error: %s", r.Method, r.URL.Path, err)

	if err := writeJSONError(w, http.StatusBadRequest, err.Error()); err != nil {
		log.Println(err)
	}
}

func (app *Application) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("not found: %s path: %s error: %s", r.Method, r.URL.Path, err)

	if err := writeJSONError(w, http.StatusNotFound, store.ErrNotFound.Error()); err != nil {
		log.Println(err)
	}
}

func (app *Application) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("concurrency update: %s path: %s error: %s", r.Method, r.URL.Path, err)

	if err := writeJSONError(w, http.StatusConflict, store.ErrConcurrentUpdate.Error()); err != nil {
		log.Println(err)
	}
}

func (app *Application) writeNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

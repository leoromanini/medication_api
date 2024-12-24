package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/medications", func(r chi.Router) {
		r.Get("/", app.medicationsList)
		r.Post("/", app.medicationsCreate)

		r.Route("/{medicationID}", func(r chi.Router) {
			r.Use(app.medicationCtx)
			r.Get("/", app.medicationGet)
			// r.Patch("/", app.medications.Update)
		})
	})

	// router := httprouter.New()
	// router.HandlerFunc(http.MethodGet, "/medications", app.medicationsLatest)

	return r
}

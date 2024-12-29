package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (app *application) routes() http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(secureHeaders)
	r.Use(jsonContentTypeHeaders)
	r.Use(httprate.LimitByIP(100, 1*time.Minute))
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(PrometheusMiddleware)

	r.Get("/", app.homePage)
	r.Route("/v1", func(r chi.Router) {
		r.Route("/medications", func(r chi.Router) {
			r.Get("/", app.medicationsList)
			r.Post("/", app.medicationCreate)

			r.Route("/{medicationID}", func(r chi.Router) {
				r.Use(app.medicationCtx)
				r.Get("/", app.medicationGet)
				r.Patch("/", app.medicationUpdate)
				r.Delete("/", app.medicationDelete)
			})
		})
	})

	r.Route("/health", func(r chi.Router) {
		r.Get("/", app.healthCheck)
	})

	r.Handle("/metrics", promhttp.Handler())

	return r
}

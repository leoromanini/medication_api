package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/leoromanini/medication_api/internal/models"
)

type contextKey string

const medicationContextKey = contextKey("medication")

func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func jsonContentTypeHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (app *application) medicationCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		medicationID, err := strconv.Atoi(chi.URLParam(r, "medicationID"))
		if err != nil || medicationID < 1 {
			app.notFound(w)
			return
		}
		medication, err := app.medications.Get(medicationID)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				app.notFound(w)
			} else {
				app.serverError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), medicationContextKey, medication)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &responseWriter{w, http.StatusOK}

		start := time.Now()

		next.ServeHTTP(ww, r)

		path := r.URL.Path
		method := r.Method
		status := ww.statusCode

		apiRequestCount.WithLabelValues(path, method, http.StatusText(status)).Inc()
		duration := time.Since(start).Seconds()
		apiRequestDuration.WithLabelValues(path, method, http.StatusText(status)).Observe(duration)
	})
}

package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/leoromanini/medication_api/internal/models"
)

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

		ctx := context.WithValue(r.Context(), "medication", medication)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

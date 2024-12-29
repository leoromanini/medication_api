package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/leoromanini/medication_api/internal/models"
)

type MedicationsRequest struct {
	*models.Medications
	validationsErrors []ValidationError
}

type MedicationsResponse struct {
	*models.Medications
}

func (rd *MedicationsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func MedicationResponse(medication *models.Medications) *MedicationsResponse {
	resp := &MedicationsResponse{Medications: medication}

	return resp
}

func (app *application) medicationsList(w http.ResponseWriter, r *http.Request) {

	medications, err := app.medications.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	medicationsRender := []render.Renderer{}
	for _, medication := range medications {
		medicationsRender = append(medicationsRender, MedicationResponse(medication))
	}

	if err := render.RenderList(w, r, medicationsRender); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value(medicationContextKey).(*models.Medications)

	if err := render.Render(w, r, MedicationResponse(medication)); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value(medicationContextKey).(*models.Medications)

	data := &MedicationsRequest{Medications: medication}
	if err := render.Bind(r, data); err != nil {
		if errors.Is(err, ErrValidation) {
			app.unprocessableEntity(w)
			if err := render.Render(w, r, ValidationsErrorResponse(data.validationsErrors)); err != nil {
				app.serverError(w, err)
				return
			}
		}
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			app.badRequestByDecode(w, extractReadableUnmarshalError(err))
			return
		}
		app.badRequest(w)
		return
	}

	medication = data.Medications
	err := app.medications.Update(medication.ID, medication.Name, medication.Dosage, medication.Form)
	if err != nil {
		app.serverError(w, err)
		return
	}

	medication, err = app.medications.Get(medication.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if err := render.Render(w, r, MedicationResponse(medication)); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationCreate(w http.ResponseWriter, r *http.Request) {
	data := &MedicationsRequest{}
	if err := render.Bind(r, data); err != nil {
		if errors.Is(err, ErrValidation) {
			app.unprocessableEntity(w)
			if err := render.Render(w, r, ValidationsErrorResponse(data.validationsErrors)); err != nil {
				app.serverError(w, err)
				return
			}
		}
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			app.badRequestByDecode(w, extractReadableUnmarshalError(err))
			return
		}
		app.badRequest(w)
		return
	}

	medication := data.Medications
	medicationID, err := app.medications.Create(medication.Name, medication.Dosage, medication.Form)
	if err != nil {
		app.serverError(w, err)
		return
	}

	medication, err = app.medications.Get(medicationID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, MedicationResponse(medication)); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationDelete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value(medicationContextKey).(*models.Medications)

	err := app.medications.Delete(medication.ID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if err := render.Render(w, r, MedicationResponse(medication)); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status": "ok"}`))
}

type HomePageResponse struct {
	Name          string            `json:"name"`
	Version       string            `json:"version"`
	Description   string            `json:"description"`
	Endpoints     map[string]string `json:"endpoints"`
	Documentation string            `json:"documentation"`
	UsageExamples string            `json:"usage_examples"`
}

func (app *application) homePage(w http.ResponseWriter, r *http.Request) {
	response := HomePageResponse{
		Name:        "Medications API",
		Version:     "0.0.1",
		Description: "This is the REST API for Medications",
		Endpoints: map[string]string{
			"/v1/medications": "CRUD operations for medications",
			"/health":         "API Healthcheck",
			"/metrics":        "API metrics for Prometheus",
		},
		Documentation: "https://github.com/leoromanini/medication_api",
		UsageExamples: "https://github.com/leoromanini/medication_api/tree/main/_examples",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		app.serverError(w, err)
	}
}

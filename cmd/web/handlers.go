package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/leoromanini/medication_api/internal/models"
)

type MedicationsRequest struct {
	*models.Medications
}

type MedicationsResponse struct {
	*models.Medications
}

func (m *MedicationsRequest) Bind(r *http.Request) error {

	if m.Medications == nil {
		return errors.New("missing required Medications fields")
	}

	return nil
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

	medications_render := []render.Renderer{}
	for _, medication := range medications {
		medications_render = append(medications_render, MedicationResponse(medication))
	}

	if err := render.RenderList(w, r, medications_render); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value("medication").(*models.Medications)

	if err := render.Render(w, r, MedicationResponse(medication)); err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) medicationUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value("medication").(*models.Medications)

	data := &MedicationsRequest{Medications: medication}
	err := render.Bind(r, data)
	if err != nil {
		app.badRequest(w)
		return
	}

	medication = data.Medications
	err = app.medications.Update(medication.ID, medication.Name, medication.Dosage, medication.Form)
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
	medication := ctx.Value("medication").(*models.Medications)

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

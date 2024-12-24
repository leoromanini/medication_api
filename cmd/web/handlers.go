package main

import (
	"fmt"
	"net/http"

	"github.com/leoromanini/medication_api/internal/models"
)

type MedicationRequest struct {
	*models.Medications
}

func (app *application) medicationsList(w http.ResponseWriter, r *http.Request) {

	medications, err := app.medications.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, medication := range medications {
		fmt.Fprintf(w, "%+v\n", medication)
	}
}

func (app *application) medicationGet(w http.ResponseWriter, r *http.Request) {
	// TODO: Maybe remove ok from here and add a recover middleware
	ctx := r.Context()
	medication, ok := ctx.Value("medication").(*models.Medications)
	if !ok {
		app.unprocessableEntity(w)
		return
	}
	fmt.Fprintf(w, "%+v\n", medication)
}

func (app *application) medicationUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	medication := ctx.Value("medication").(*models.Medications)

	data := &models.Medication{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	article = data.Article
	dbUpdateArticle(article.ID, article)

	render.Render(w, r, NewArticleResponse(article))

}

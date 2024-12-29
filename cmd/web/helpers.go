package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) notFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "{\"error\": \"Medication ID Not Found\"}")
}

func (app *application) unprocessableEntity(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnprocessableEntity)
}

func (app *application) badRequestByDecode(w http.ResponseWriter, m string) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintln(w, "{\"error\": \""+m+"\"}")
}

func (app *application) badRequest(w http.ResponseWriter) {
	http.Error(w, "", http.StatusBadRequest)
}

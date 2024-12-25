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

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// TODO: Return that for validation errors
// func (app *application) unprocessableEntity(w http.ResponseWriter) {
// 	app.clientError(w, http.StatusUnprocessableEntity)
// }

func (app *application) badRequest(w http.ResponseWriter) {
	app.clientError(w, http.StatusBadRequest)
}

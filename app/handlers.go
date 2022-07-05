package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) Library(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := app.renderTemplate(w, r, "library", nil)
	if err != nil {
		app.errorLog.Println(err)
	}
}

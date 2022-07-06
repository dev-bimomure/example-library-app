package main

import (
	"net/http"

	"github.com/dev-bimomure/example-library-app/app/models"
	"github.com/julienschmidt/httprouter"
)

func (app *application) Library(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	books, err := app.DB.GetAllBooks()
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	booksData := models.Book{
		ID:        books.ID,
		Name:      books.Name,
		Publisher: books.Publisher,
		Date:      books.Date,
	}

	data := make(map[string]interface{})
	data["books"] = booksData

	err = app.renderTemplate(w, r, "library", &templateData{
		Data: data,
	})
	if err != nil {
		app.errorLog.Println(err)
	}
}

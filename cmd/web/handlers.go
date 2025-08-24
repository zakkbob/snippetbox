package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/zakkbob/snippetbox/internal/models"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, statusCode int, page string, data any) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(statusCode)
	buf.WriteTo(w)
}

// GET '/' - Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go!")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Snippets: snippets,
	}

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)
}

// GET '/snippet/view/{id}' - View a snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Return 404 if the id is not an integer above 0
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	data := templateData{
		Snippet: snippet,
	}

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

// GET '/snippet/create' - Create a snippet?
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "_wow_, you just tried to create a snippet using a GET request. Try POST next time!")
}

// POST '/snippet/create' - Create a snippet, but with POST this time!
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	var (
		title   = "Test snippet"
		content = "Hello person reading this"
		expires = 7
	)

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

// ----------------------------------------------------- //

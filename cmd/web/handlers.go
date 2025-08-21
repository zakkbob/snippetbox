package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/zakkbob/snippetbox/internal/models"
)

// ---- Handlers which serve their respective pages ---- //

// GET '/' - Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go!")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// // Define the templates to be parsed, order doesn't matter as we are using ExecuteTemplate
	// files := []string{
	// "./ui/html/pages/base.tmpl.html",
	// "./ui/html/partials/nav.tpml.html",
	// "./ui/html/pages/home.tmpl.html",
	// }
	//
	// // Add the template files into a template set. Handle error appropriately if it occurs
	// ts, err := template.ParseFiles(files...)
	// if err != nil {
	// app.serverError(w, r, err)
	// return
	// }
	//
	// // Write the content of "base" template to the Response Body
	// err = ts.ExecuteTemplate(w, "base", nil)
	// if err != nil {
	// app.serverError(w, r, err)
	// }
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

	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/nav.tpml.html",
		"./ui/html/pages/view.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Snippet: snippet,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, r, err)
	}
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

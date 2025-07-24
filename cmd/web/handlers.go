package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"
)

// ---- Handlers which serve their respective pages ---- //

// GET '/' - Home page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go!")

	// Define the templates to be parsed, order doesn't matter as we are using ExecuteTemplate
	files := []string{
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/partials/nav.tpml.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// Add the template files into a template set. Handle error appropriately if it occurs
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Write the content of "base" template to the Response Body
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// GET '/snippet/view/{id}' - View a snippet
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// Return 404 if the id is not an integer above 0
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Wow, you just found snippet %d!", id)
}

// GET '/snippet/create' - Create a snippet?
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "_wow_, you just tried to create a snippet using a GET request. Try POST next time!")
}

// POST '/snippet/create' - Create a snippet, but with POST this time!
func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Wow, you just created a snippet using a POST request!"))
}

// ----------------------------------------------------- //

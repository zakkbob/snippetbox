package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// ---- Handlers which serve their respective pages ---- //

// GET '/' - Home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// GET '/snippet/view/{id}' - View a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Retur 404 if the id is not an integer above 0
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Wow, you just found snippet %d!", id)
	w.Write([]byte(msg))
}

// GET '/snippet/create' - Create a snippet?
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("_wow_, you just tried to create a snippet using a GET request. Try POST next time!"))
}

// POST '/snippet/create' - Create a snippet, but with POST this time!
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wow, you just created a snippet using a POST request!"))
}

// ----------------------------------------------------- //

func main() {
	// Initialise a 'servemux', this is where route handlers will be registered
	mux := http.NewServeMux()

	// Register the handlers for the directories. '/{$}' is used so that home is no longer a catch-all - a 404 will be returned instead
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView) // {id} is a wildcard route pattern, it'll match any non-empty value in that segment. Also, use 'id' instead of 'snippetID' as this avoids 'stutter'

	log.Print("starting server on :4000")

	// Start a new web server on port 4000, using the servemux we just created
	// Then, log an error if we get one, Fatal will terminate the program
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}

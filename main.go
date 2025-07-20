package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// ---- Handlers which serve their respective pages ---- //

// GET '/' - Home page
func home(w http.ResponseWriter, r *http.Request) {
	// Add custom header 'Server: Go!'
	w.Header().Add("Server", "Go!")
	w.Write([]byte("Hello World!"))
}

// GET '/snippet/view/{id}' - View a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Return 404 if the id is not an integer above 0
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// http.ResponseWriter satisfies the io.Writer interface, so can be used directly with Fprintf!
	fmt.Fprintf(w, "Wow, you just found snippet %d!", id)
}

// GET '/snippet/create' - Create a snippet?
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// http.ResponseWriter satisfies the io.Writer interface, so can be using with io.WriteString!
	io.WriteString(w, "_wow_, you just tried to create a snippet using a GET request. Try POST next time!")
}

// POST '/snippet/create' - Create a snippet, but with POST this time!
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Send a 201 Created status code rather than 200 OK
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Wow, you just created a snippet using a POST request!"))
}

// ----------------------------------------------------- //

func main() {
	// Initialise a 'servemux', this is where route handlers will be registered
	mux := http.NewServeMux()

	// Register the handlers for the directories and specify the HTTP method. '/{$}' is used so that home is no longer a catch-all - a 404 will be returned instead
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

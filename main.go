package main

import (
	"log"
	"net/http"
)

// ---- Handlers which serve their respective pages ---- //

// '/' - Home page
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// '/snippet/view' - View a snippet
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wow, you just found a snippet!"))
}

// '/snippet/create' - Create a snippet?
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wow, you just created a snippet!"))
}

// ----------------------------------------------------- //

func main() {
	// Initialise a 'servemux', this is where route handlers will be registered
	mux := http.NewServeMux()

	// Register the handlers for the directories. '/{$}' is used so that home is no longer a catch-all - a 404 will be returned instead
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on :4000")

	// Start a new web server on port 4000, using the servemux we just created
	// Then, log an error if we get one, Fatal will terminate the program
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}

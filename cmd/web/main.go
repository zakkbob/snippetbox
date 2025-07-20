package main

import (
	"log"
	"net/http"
)

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

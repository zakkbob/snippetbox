package main

import (
	"log"
	"net/http"
)

// Function which serves the 'home' page - '/'
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	// Initialise a 'servemux', I don't know what this means yet
	mux := http.NewServeMux()
	// Register the home function as the handler for the root directory. ('/' is also a catch-all so it handles ALL directories)
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// Start a new web server on port 4000, using the servemux we just created
	// Then, log an error if we get one, Fatal will terminate the program
	err := http.ListenAndServe("localhost:4000", mux)
	log.Fatal(err)
}

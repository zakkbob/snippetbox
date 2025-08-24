package main

import (
	"net/http"
)

func (a *application) routes() http.Handler {
	// Initialise a 'servemux', this is where route handlers will be registered
	mux := http.NewServeMux()

	// Create a file server, to serve files out of ui/static
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register fileServer as the handler for all paths starting with '/static/'
	// Remove the 'static/' prefix before it reaches the fileserver
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the handlers for the directories and specify the HTTP method. '/{$}' is used so that home is no longer a catch-all - a 404 will be returned instead
	mux.HandleFunc("GET /{$}", a.home)
	mux.HandleFunc("POST /snippet/create", a.snippetCreatePost)
	mux.HandleFunc("GET /snippet/view/{id}", a.snippetView) // {id} is a wildcard route pattern, it'll match any non-empty value in that segment. Also, use 'id' instead of 'snippetID' as this avoids 'stutter'

	// mux.HandleFunc is syntactic sugar for this, so we can just use this directly instead
	mux.Handle("GET /snippet/create", http.HandlerFunc(a.snippetCreate))

	return a.logRequest(commonHeaders(mux))
}

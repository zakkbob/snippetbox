package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Holds application wide dependencies, allowing for dependency injection
type application struct {
	logger *slog.Logger
}

func main() {
	// Initialise a new structured logger, along with its handler
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Get network address from flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// Parse the command-line flags. Otherwise the default value will remain. This is why flag.String() returns a pointer
	flag.Parse()

	app := &application{
		logger: logger,
	}

	// Initialise a 'servemux', this is where route handlers will be registered
	mux := http.NewServeMux()

	// Create a file server, to serve files out of ui/static
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register fileServer as the handler for all paths starting with '/static/'
	// Remove the 'static/' prefix before it reaches the fileserver
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the handlers for the directories and specify the HTTP method. '/{$}' is used so that home is no longer a catch-all - a 404 will be returned instead
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView) // {id} is a wildcard route pattern, it'll match any non-empty value in that segment. Also, use 'id' instead of 'snippetID' as this avoids 'stutter'

	// mux.HandleFunc is syntactic sugar for this, so we can just use this directly instead
	mux.Handle("GET /snippet/create", http.HandlerFunc(app.snippetCreate))

	// Structured log, specifies a key-value pair
	logger.Info("starting server", "address", *addr)

	// Start a new web server on configured port, using the servemux we just created
	// Then, log an error if we get one, slog doesn't have an equivalent to Fatal(), so calling os.Exit(1) is necess:ary
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

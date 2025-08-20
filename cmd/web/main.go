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
	// Get network address from flag
	addr := flag.String("addr", ":4000", "HTTP Network Address")

	// Parse the command-line flags. Otherwise the default value will remain. This is why flag.String() returns a pointer
	flag.Parse()

	// Initialise a new structured logger, along with its handler
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "address", *addr)

	// Start a new web server on configured port, using the servemux we just created
	// Then, log an error if we get one, slog doesn't have an equivalent to Fatal(), so calling os.Exit(1) is necess:ary
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

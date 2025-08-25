package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register fileServer as the handler for all paths starting with '/static/'
	// Remove the 'static/' prefix before it reaches the fileserver
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	handleSessions := alice.New(app.sessionManager.LoadAndSave)

	mux.Handle("GET /{$}", handleSessions.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", handleSessions.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", handleSessions.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", handleSessions.ThenFunc(app.userSignupPost))
	mux.Handle("GET /user/login", handleSessions.ThenFunc(app.userLogin))
	mux.Handle("POST /user/login", handleSessions.ThenFunc(app.userLoginPost))

	requireAuth := handleSessions.Append(app.requireAuthentication)

	mux.Handle("GET /snippet/create", requireAuth.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", requireAuth.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", requireAuth.ThenFunc(app.userLogoutPost))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}

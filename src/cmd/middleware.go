package main

import "net/http"

func (app *Application) AuthorizationMiddleware(next http.Handler) http.Handler {
	// Do auth checks
	return next
}

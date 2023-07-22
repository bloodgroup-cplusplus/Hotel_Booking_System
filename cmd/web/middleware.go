package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"

)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// No surf adds CSRF protection to all post request

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// middleware to remember states using session tokens it loads and saves thes ession on every request

func SessionLoad(next http.Handler) http.Handler {
	// Load and Save provides middleware that automatically loads and saves session
	// Data for the current request, and communicates the session token to and from
	// the client in a cookie.
	return session.LoadAndSave((next))
}
abcdefghij
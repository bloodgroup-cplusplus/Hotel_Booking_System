package main

import (
	"net/http"

	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/config"
	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	// create a new http handler
	// using mux which is known as multiplexer

	// create a multiplexer
	// which is a http handler

	//mux := pat.New()

	// route to our home function
	// mux has methods built it in it
	//mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// let's use the pat router

	// let's create a new chi mux

	mux := chi.NewRouter()
	// how to use middlewares to recover from a panic that program goes
	mux.Use(middleware.Recoverer)
	// let's use more middlewares in upcoming lectures
	// let's write to console
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// writing custom middlewares when anyone visits the website
	mux.Use(WriteToConsole)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// create a file server to so that we find our public /static files to serve the content
	fileServer := http.FileServer(http.Dir("./static/"))
	// use the file server in our multiplexers
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}

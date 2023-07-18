package main

import (
	"net/http"

	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/config"
	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/handlers"
	"github.com/go-chi/chi"
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
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}

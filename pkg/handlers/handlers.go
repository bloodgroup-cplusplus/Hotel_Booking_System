package handlers

import (
	"net/http"

	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/config"
	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// New Repo creates the new rpository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{

		App: a,
	}
}

// New Handle33rs sets teh repositeor y of the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "about.page.tmpl")
}

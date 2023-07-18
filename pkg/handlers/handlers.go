package handlers

import (
	"net/http"

	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/config"
	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/models"
	"github.com/bloodgroup-cplusplus/Hotel_Booking_System/pkg/render"
)

// Template Data holds data set from handlers to templates

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
	remoteIP := r.RemoteAddr
	// everytime somebody hits the home page
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	// perform some logic
	// send the data to the template
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello , again."
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap})
}

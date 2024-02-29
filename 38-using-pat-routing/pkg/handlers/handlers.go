package handlers

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bookings-udemy/pkg/render"
	"net/http"
)

// Following the repository pattern here (look at screenshots for more information)

// Create a new type Repository which is a struct, and it has a field App of type *config.AppConfig
type Repository struct {
	App *config.AppConfig
}

// Declaring a variable Repo of type *Repository (pointer to Repository)
// Repo the repository used, by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	config := &Repository{
		App: a,
	}

	return config
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Remote ip address
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	// Perform some business logic
	templateData := &models.TemplateData{
		StringMap: map[string]string{
			"test":      "Hello, this is a test",
			"remote_ip": remoteIP,
		},
	}

	// Send the data to the template
	render.RenderTemplate(w, "about.page.tmpl", templateData)
}

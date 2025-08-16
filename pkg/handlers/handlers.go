package handlers

import (
	"github.com/nowa75/bookings/pkg/config"
	"github.com/nowa75/bookings/pkg/models"
	"github.com/nowa75/bookings/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.TemplateRender(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perf some logic
	stringMap := make(map[string]string)
	stringMap["About"] = "Welcome to the About Page"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.TemplateRender(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}

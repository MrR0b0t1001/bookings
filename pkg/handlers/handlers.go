package handlers

import (
	"net/http"

	"github.com/MrR0b0t1001/bookings/models"
	"github.com/MrR0b0t1001/bookings/pkg/config"
	"github.com/MrR0b0t1001/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// Repo variable used by the handlers
var Repo *Repository

// Creates our new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets our Repo variable to r
func NewHandlers(r *Repository) {
	Repo = r
}

func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr // holds the IP address of the request

	re.App.SessionManager.Put(r.Context(), "remote_ip", remoteIP) // We store it

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIP := re.App.SessionManager.GetString(
		r.Context(),
		"remote_ip",
	) // we retrieve the value of the key

	stringMap := map[string]string{}
	stringMap["test"] = "Good Bye"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

package handlers

import (
	"github.com/tnthanh47/GoFirstProject/pkg/config"
	"github.com/tnthanh47/GoFirstProject/pkg/models"
	"github.com/tnthanh47/GoFirstProject/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, request *http.Request) {

	//Perform some logic
	strMap := map[string]string{}
	strMap["test"] = "hello"
	render.RenderTemplate(
		w, "home.page.html", &models.TemplateData{
			MapString: strMap,
		},
	)
}

func (m *Repository) About(w http.ResponseWriter, req *http.Request) {
	strMap := map[string]string{}
	strMap["test"] = "hello"
	render.RenderTemplate(
		w, "about.page.html", &models.TemplateData{
			strMap,
		},
	)
}

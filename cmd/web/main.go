package main

import (
	"fmt"
	"github.com/tnthanh47/GoFirstProject/pkg/config"
	"github.com/tnthanh47/GoFirstProject/pkg/handlers"
	"github.com/tnthanh47/GoFirstProject/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var appConfig config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	appConfig.TemplateCache = tc
	appConfig.UseCache = false

	render.NewTemplateCache(&appConfig)
	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandler(repo)

	http.HandleFunc("/Home", handlers.Repo.Home)
	http.HandleFunc("/About", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Start Application listening to Port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}

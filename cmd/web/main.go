package main

import (
	"log"
	"net/http"

	"github.com/ankitdas09/gowebapp/cmd/pkg/config"
	"github.com/ankitdas09/gowebapp/cmd/pkg/handlers"
	"github.com/ankitdas09/gowebapp/cmd/pkg/render"
)

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	// http.ListenAndServe(":8080", nil)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/ankitdas09/gowebapp/cmd/pkg/config"
	"github.com/ankitdas09/gowebapp/cmd/pkg/handlers"
	"github.com/ankitdas09/gowebapp/cmd/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {
	tc, err := render.CreateTemplateCache()

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.InProduction = false

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
	log.Println("High performance GO server on port 8080")
	err = srv.ListenAndServe()
	log.Fatal(err)
}

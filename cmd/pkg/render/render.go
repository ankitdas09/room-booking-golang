package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ankitdas09/gowebapp/cmd/pkg/config"
	"github.com/ankitdas09/gowebapp/cmd/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {
	// get the cache from application config
	// tc, err := CreateTemlateCache()
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get template from cache
	tmpl, ok := tc[t]
	if !ok {
		log.Fatal("Could not get template from cache")
	}

	// render the template

	// i. check if parsed data is correct
	buff := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := tmpl.Execute(buff, td)
	if err != nil {
		log.Println(err)
	}

	// ii. render the template
	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return tCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return tCache, err
		}
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return tCache, err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return tCache, err
			}
		}
		tCache[name] = ts
	}
	return tCache, nil
}

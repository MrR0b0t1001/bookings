package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/MrR0b0t1001/bookings/models"
	"github.com/MrR0b0t1001/bookings/pkg/config"
)

// create app variable that holds our config setup
var app *config.AppConfig

// Initialize said variable app
func NewTemplate(a *config.AppConfig) {
	app = a
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// Initializing myCache to {}
	myCache := map[string]*template.Template{}

	// Get all files that end in .page.tmpl
	// GLOB GLOB ALL FILES
	pages, err := filepath.Glob("./Templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Range through the pages ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(
			page,
		) // Base return the last part of the filepath we got from Glob
		// ts contains our parsed templates
		ts, err := template.New(name).
			ParseFiles(page)
			// I create a new template that ends in .page.tmpl and Parse it
		if err != nil {
			return myCache, err
		}

		// Looking for all *.layout.tmpl files
		layouts, err := filepath.Glob("./Templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layouts) > 0 {
			// We need to parse all files that end in *.layout.tmpl
			// Because they may be needed by the *.page.tmpl
			// GLOB GLOB && Parse ALL files
			ts, err = ts.ParseGlob("./Templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// We add all parsed templates from files *.page.tmpl && *.layout.tmpl
		// We add it into a map (our cache) whose keys are the file names e.g home.page.tmpl or base.layout.tmpl
		myCache[name] = ts

	}

	return myCache, nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var (
		templateCache map[string]*template.Template
		err           error
	)
	// Retrieving templateCache

	// if templateCache exists use it
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		// else create it
		templateCache, err = CreateTemplateCache()
	}
	// For now that means we are gonna get the templates for home.page.html and about.page.html
	// parsedTemplate contains the parsed template if the string tmpl matches an entry in our Cache
	parsedTemplate, ok := templateCache[tmpl]
	if !ok {
		log.Fatalf("Could not retrieve template with key : %v", tmpl)
	}
	// Taking the buffer approach to be able to pinpoint unexpected errors more efficiently
	buf := new(bytes.Buffer)
	err = parsedTemplate.Execute(buf, td)
	// If an error occurs then we can pinpoint its location from inside the templateCache map
	if err != nil {
		log.Println(err)
	}

	// Render the Template
	_, err = buf.WriteTo(w)

	if err != nil {
		log.Println(err)
	}
}

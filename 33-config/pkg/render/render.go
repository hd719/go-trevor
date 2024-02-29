package render

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/models"
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

// Sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {

	var tc map[string]*template.Template

	// Read from disk every time in production
	if app.UseCache {
		// Instead the template cache is created once in main.go and then passed to the NewTemplates function
		tc = app.TemplateCache
	} else {
		// Create the template cache this is development mode
		tc, _ = CreateTemplateCache()
	}

	// Old way of doing creating template cache, this is not good because it will create a new template cache every time a template is rendered
	// tc, err := createTemplateCache()

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Hold bytes in memory
	buf := new(bytes.Buffer)

	// Add default data to the template
	td := AddDefaultData(data)

	// Execute the template (error handling is not shown)
	_ = t.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

// Returns a parsed template that can be executed
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// Get all of the files named *page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// Loop through the pages
	for _, page := range pages {
		// Get the filename
		name := filepath.Base(page)

		// Create a template set and then call the ParseFiles function (similar to line 19)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// This will return the layout template
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		// If there are any layout templates
		if len(matches) > 0 {
			// Combines the template that is returned on line 42 with the layout template (which is base.layout.tmpl)
			_, err := ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		// Add the template set to the cache
		myCache[name] = ts
	}

	return myCache, nil
}

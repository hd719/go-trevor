package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// Create template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	// Could also be written as:
	// else if ok {
	// 	t.Execute(w, nil)
	// } else {
	// 	log.Fatal("Could not get template from template cache")
	// }

	// Hold bytes in memory
	buf := new(bytes.Buffer)

	// Execute the template (error handling is not shown)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println("Error writing template to browser", err)
	}
}

// Returns a parsed template that can be executed
func createTemplateCache() (map[string]*template.Template, error) {
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

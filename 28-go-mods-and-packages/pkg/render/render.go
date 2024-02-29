package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders a template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	// This is quite inefficient, as it reads the file every time a page is rendered
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// Check to see if we already have the template in the cache
	_, inMap := tc[t]

	if !inMap {
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println("error creating template cache", err)
			return
		}
	} else {
		// we have template in cache
		log.Println("using template from cache")
	}

	// get the template from the cache
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println("error executing template", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the template files
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add the template to the cache
	tc[t] = tmpl

	return nil
}

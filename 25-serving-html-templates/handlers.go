package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Home is the handler for the home page
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About is the handler for the about page
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")

}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/", tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

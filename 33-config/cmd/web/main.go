package main

import (
	"bookings-udemy/pkg/config"
	"bookings-udemy/pkg/handlers"
	"bookings-udemy/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache() // Create the template cache map data structure with key value pairs of string and *template.Template
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	// Assign the template cache to the app config (which will be used throughout our application)
	app.TemplateCache = tc
	app.UseCache = false

	// Create a new repository and pass a reference to the app config
	repo := handlers.NewRepo(&app)
	// Create a new handlers and pass a reference
	handlers.NewHandlers(repo)

	// Pass a reference to the app config to the NewTemplates function in render package
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// By passing a reference to the app variable (&app) instead of its value, any modifications made within these functions to the app variable will be reflected in the original variable declared in the main function. This ensures that all components of the application share the same configuration.

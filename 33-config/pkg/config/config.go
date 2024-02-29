package config

import "html/template"

// AppConfig holds the application configuration, which is initialized in main.go
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
}

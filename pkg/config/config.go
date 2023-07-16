package config

import (
	"html/template"
	"log"
)

// AppConfig holds the applicaiton config
type AppConfig struct {
	// structure that holds some type of structure
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	// write logs to files databases etc
	// it is accessible to every part that is accessible to app configuraiton so useful
}

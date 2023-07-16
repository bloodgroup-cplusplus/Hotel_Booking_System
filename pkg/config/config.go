package config

import "html/template"

// AppConfig holds the applicaiton config
type AppConfig struct {
	// structure that holds some type of structure
	TemplateCache map[string]*template.Template
}

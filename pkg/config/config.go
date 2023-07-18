package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the applicaiton config
type AppConfig struct {
	// structure that holds some type of structure
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	// write logs to files databases etc
	// it is accessible to every part that is accessible to app configuraiton so useful
	// we use one more value called inProduction to know whether we are in development or whether we are in
	// production
	InProduction bool
	Session      *scs.SessionManager
}

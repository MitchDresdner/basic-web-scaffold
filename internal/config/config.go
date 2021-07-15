package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	Domain        string
	Identifier    string
	InfoLog       *log.Logger
	InProduction  bool
	PreferenceMap map[string]string
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
	UseCache      bool
	Version       string
}

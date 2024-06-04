package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)
// AppConfig Holds the App configuration
type AppConfig struct {
	UseCache       bool // False : Reads templates from disk every time, True : Will use existing cache and not read
	TemplateCache  map[string]*template.Template
	InProduction   bool // False indicates still in developer mode
	SessionManager *scs.SessionManager
}

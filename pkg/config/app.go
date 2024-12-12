package config

import (
	"github.com/rs/zerolog/log"
)

// App holds the App configuration
type App struct {
	Host string
	Port int
}

var app = &App{}

// AppCfg returns the default App configuration
func AppCfg() *App {
	return app
}

// LoadApp loads App configuration
func LoadApp() {
	log.Info().Msgf("PORT: 3000")
	app.Port = 3000
}

package config

import (
	"github.com/amy911/env911/app"
)

// Config describes the app configuration.
type Config struct {
	app *app.Apper
	map_ map[string]interface{}
}

// App gets the app being configured.
func (config *Config) App() *app.Apper {
	return config.app
}

// Map gets the map of keys to configured values.
func (config *Config) Map() map[string]interface{} {
	return config.map_
}

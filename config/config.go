package config

import (
	"github.com/amy911/env911/app"
)

// Config describes the app configuration.
type Config struct {
	app *app.Apper
	map_ map[string]interface{}
}

// New creates a new Config.
func New(app *app.Apper, args ...interface{}) *Config {
	return new(Config).Init(app, args...)
}

// App gets the app being configured.
func (config *Config) App() *app.Apper {
	return config.app
}

// Init initializes a Config.
func (config *Config) Init(app *app.Apper, args ...interface{}) *Config {
	config.app = app
	config.map_ = make(map[string]interface{})
	return config
}

// Map gets the map of keys to configured values.
func (config *Config) Map() map[string]interface{} {
	return config.map_
}

package config

import (
	"github.com/amy911/env911/app"
)

// Config describes the app configuration.
type Config struct {
	// App is the app being configured.
	App *app.Apper

	// Map is the map of keys to configured values.
	Map map[string]interface{}
}

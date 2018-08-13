package config

import (
	"github.com/amy911/amy911/onfail"
	"github.com/amy911/env911/app"
)

// Configger is the interface for app configuration.
type Configger interface {

	// App gets the app being configured.
	App() app.Apper

	// Env gets the map of keys to values configured via environment variables.
	Env() map[string]interface{}

	// Local gets the map of keys to values configured locally.
	Local() map[string]interface{}

	// System gets the map of keys to values configured system-wide.
	System() map[string]interface{}

}

// Get gets the value associated with key from the configuration sources
func (config Configger) Get(key string) interface{} {
	if v, ok := config.Env(); ok {
		return v
	}
	if v, ok := config.Local(); ok {
		return v
	}
	if v, ok := config.System(); ok {
		return v
	}
	return nil
}

// Load loads mappings from the several sources.
func (config Configger) Load(onFail ...onfail.OnFail) error {
	config.LoadConfigFile(config.System(), config.App().SystemConfigFile(), ...onFail)
	config.LoadConfigFile(config.Local(), config.App().LocalConfigFile(), ...onFail)
	config.LoadEnv(config.Env(), ...onFail)
}

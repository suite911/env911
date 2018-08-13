package config

import (
	"github.com/amy911/amy911/onfail"
	"github.com/amy911/env911/app"
)

// Configger is the interface for app configuration.
type Configger interface {

	// App gets the app being configured.
	App() app.Apper

	// AutoBind gets whether or not to use automatic environment variable binding.
	AutoBind() bool

	// Bind binds the an environment key.
	Bind(string) Configger

	// Env gets the map of keys to values configured via environment variables.
	Env() map[string]interface{}

	// FlagSet gets the flag set used for configuration.
	FlagSet() Flagger

	// Get gets the value associated with key from the configuration sources
	Get(key string) interface{}

	// Load loads mappings from the several sources.
	Load()

	// LoadEnv loads mappings from environment variables.
	LoadEnv()

	// LoadLocal loads mappings from the local configuration file.
	LoadLocal() error

	// LoadSystem loads mappings from the system configuration file.
	LoadSystem() error

	// Local gets the map of keys to values configured locally.
	Local() map[string]interface{}

	// Prefix gets the environment variable prefix.
	Prefix() string

	// SetAutoBind sets whether or not to use automatic environment variable binding.
	SetAutoBind(bool) Configger

	// SetPrefix sets the environment variable prefix.
	SetPrefix(string) Configger

	// System gets the map of keys to values configured system-wide.
	System() map[string]interface{}

}

func (config Configger) Bool(name string, value bool, usage string) {
	config.FlagSet().Bool(name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) BoolP(name, shorthand string, value bool, usage string) {
	config.FlagSet().BoolP(name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) BoolVar(p *bool, name string, value bool, usage string) {
	config.FlagSet().BoolVar(p, name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	config.FlagSet().BoolVarP(p, name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) Count(name string, by int, usage string) {
	config.FlagSet().Count(name, by, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) CountP(name, shorthand string, by int, usage string) {
	config.FlagSet().CountP(name, shorthand, by, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) CountVar(p *int, name string, by int, usage string) {
	config.FlagSet().CountVar(p, name, by, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) CountVarP(p *int, name, shorthand string, by int, usage string) {
	config.FlagSet().CountVarP(p, name, shorthand, by, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) Int(name string, value int, usage string) {
	config.FlagSet().Int(name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) IntP(name, shorthand string, value int, usage string) {
	config.FlagSet().IntP(name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) IntVar(p *int, name string, value int, usage string) {
	config.FlagSet().IntVar(p, name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) IntVarP(p *int, name, shorthand string, value int, usage string) {
	config.FlagSet().IntVarP(p, name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) String(name, value, usage string) {
	config.FlagSet().String(name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) StringP(name, shorthand, value string, usage string) {
	config.FlagSet().StringP(name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) StringVar(p *string, name, value, usage string) {
	config.FlagSet().StringVar(p, name, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

func (config Configger) StringVarP(p *string, name, shorthand, value string, usage string) {
	config.FlagSet().StringVarP(p, name, shorthand, value, usage)
	if config.AutoBind() {
		config.Bind(name)
	}
}

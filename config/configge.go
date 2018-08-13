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

	// Flag gets the Flagger used for configuration.
	Flag() Flagger

	// Local gets the map of keys to values configured locally.
	Local() map[string]interface{}

	// System gets the map of keys to values configured system-wide.
	System() map[string]interface{}

}

func (config Configger) Bool(name string, value bool, usage string) {
	config.Flag().Bool(name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) BoolP(name, shorthand string, value bool, usage string) {
	config.Flag().BoolP(name, shorthand value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) BoolVar(p *bool, name string, value bool, usage string) {
	config.Flag().BoolVar(p, name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	config.Flag().BoolVarP(p, name, shorthand, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) Count(name string, value bool, usage string) {
	config.Flag().Count(name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) CountP(name, shorthand string, value bool, usage string) {
	config.Flag().CountP(name, shorthand value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) CountVar(p *int, name string, value bool, usage string) {
	config.Flag().CountVar(p, name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) CountVarP(p *int, name, shorthand string, value bool, usage string) {
	config.Flag().CountVarP(p, name, shorthand, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
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

func (config Configger) Int(name string, value int, usage string) {
	config.Flag().Int(name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) IntP(name, shorthand string, value int, usage string) {
	config.Flag().IntP(name, shorthand value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) IntVar(p *int, name string, value int, usage string) {
	config.Flag().IntVar(p, name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) IntVarP(p *int, name, shorthand string, value int, usage string) {
	config.Flag().IntVarP(p, name, shorthand, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

// Load loads mappings from the several sources.
func (config Configger) Load(onFail ...onfail.OnFail) error {
	config.LoadConfigFile(config.System(), config.App().SystemConfigFile(), ...onFail)
	config.LoadConfigFile(config.Local(), config.App().LocalConfigFile(), ...onFail)
	config.LoadEnv(config.Env(), ...onFail)
}

func (config Configger) String(name, value, usage string) {
	config.Flag().String(name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) StringP(name, shorthand, value string, usage string) {
	config.Flag().StringP(name, shorthand value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) StringVar(p *string, name, value, usage string) {
	config.Flag().StringVar(p, name, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

func (config Configger) StringVarP(p *string, name, shorthand, value string, usage string) {
	config.Flag().StringVarP(p, name, shorthand, value, usage)
	if config.AutomaticEnv() {
		config.AddKey(name)
	}
}

package config

import (
	"github.com/suite911/env911/app"

	"github.com/suite911/flag911/flag"
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
	FlagSet() *flag.FlagSet

	// Get gets the value associated with key from the configuration sources
	Get(key string) interface{}

	// Load loads mappings from the several sources.
	Load() Configger

	// LoadAndParse loads mappings from the several sources and parses flags on the command line.
	LoadAndParse()

	// LoadEnv loads mappings from environment variables.
	LoadEnv()

	// LoadLocal loads mappings from the local configuration file.
	LoadLocal() error

	// LoadSystem loads mappings from the system configuration file.
	LoadSystem() error

	// Local gets the map of keys to values configured locally.
	Local() map[string]interface{}

	// Parse parses flags on the command line.
	Parse(arguments []string) error

	// Prefix gets the environment variable prefix.
	Prefix() string

	// QueueStore queues a single key-value pair for storage into the local configuration file upon the next call to Save.
	QueueStore(string, interface{}) Configger

	// Save atomically saves changes to the local configuration file.
	Save() error

	// SetAutoBind sets whether or not to use automatic environment variable binding.
	SetAutoBind(bool) Configger

	// SetPrefix sets the environment variable prefix.
	SetPrefix(string) Configger

	// System gets the map of keys to values configured system-wide.
	System() map[string]interface{}

}

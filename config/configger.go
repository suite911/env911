package config

import "github.com/amy911/env911/app"

// Configger is the interface for app configuration.
type Configger interface {

	// App gets the app being configured.
	App() app.Apper

	// AutoBind gets whether or not to use automatic environment variable binding.
	AutoBind() bool

	// Bind binds the an environment key.
	Bind(string) Configger

	// Forward the call to the Flagger.
	Bool(name string, value bool, usage string) *bool

	// Forward the call to the Flagger.
	BoolP(name, shorthand string, value bool, usage string) *bool

	// Forward the call to the Flagger.
	BoolVar(p *bool, name string, value bool, usage string)

	// Forward the call to the Flagger.
	BoolVarP(p *bool, name, shorthand string, value bool, usage string)

	// Forward the call to the Flagger.
	Count(name string, by int, usage string) *int

	// Forward the call to the Flagger.
	CountP(name, shorthand string, by int, usage string) *int

	// Forward the call to the Flagger.
	CountVar(p *int, name string, by int, usage string)

	// Forward the call to the Flagger.
	CountVarP(p *int, name, shorthand string, by int, usage string)

	// Env gets the map of keys to values configured via environment variables.
	Env() map[string]interface{}

	// FlagSet gets the flag set used for configuration.
	FlagSet() Flagger

	// Get gets the value associated with key from the configuration sources
	Get(key string) interface{}

	// Forward the call to the Flagger.
	Int(name string, value int, usage string) *int

	// Forward the call to the Flagger.
	IntP(name, shorthand string, value int, usage string) *int

	// Forward the call to the Flagger.
	IntVar(p *int, name string, value int, usage string)

	// Forward the call to the Flagger.
	IntVarP(p *int, name, shorthand string, value int, usage string)

	// Load loads mappings from the several sources.
	Load() Configger

	// Load and Parse.
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

	// Forward the call to the Flagger.
	String(name, value, usage string) *string

	// Forward the call to the Flagger.
	StringP(name, shorthand, value string, usage string) *string

	// Forward the call to the Flagger.
	StringVar(p *string, name, value, usage string)

	// Forward the call to the Flagger.
	StringVarP(p *string, name, shorthand, value string, usage string)

	// System gets the map of keys to values configured system-wide.
	System() map[string]interface{}

}

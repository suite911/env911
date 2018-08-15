package config

import (
	"sync"

	"github.com/suite911/env911/app"

	"github.com/suite911/error911/onfail"
	"github.com/suite911/flag911/flag"
)

// Init initializes the configuration for this app.
func Init(config Configger, onFail ...onfail.OnFail) {
	mutex.Lock(); defer mutex.Unlock()
	if self != nil {
		onfail.Fail("Double init", self, onfail.Panic, onFail)
		return
	}
	self = config
}

// IsInit reports whether or not the configuration for this app has been initialized.
func IsInit() bool {
	return self != nil
}

// App gets the app being configured.
func App() app.Apper {
	return self.App()
}

// AutoBind gets whether or not to use automatic environment variable binding.
func AutoBind() bool {
	return self.AutoBind()
}

// Bind binds the an environment key.
func Bind(key string) Configger {
	return self.Bind(key)
}


// Env gets the map of keys to values configured via environment variables.
func Env() map[string]interface{} {
	return self.Env()
}

// FlagSet gets the flag set used for configuration.
func FlagSet() *flag.FlagSet {
	return self.FlagSet()
}

// Get gets the value associated with key from the configuration sources
func Get(key string) interface{} {
	return self.Get(key)
}

// Load loads mappings from the several sources.
func Load() Configger {
	return self.Load()
}

// LoadAndParse loads mappings from the several sources and parses flags on the command line.
func LoadAndParse() {
	self.LoadAndParse()
}

// LoadEnv loads mappings from environment variables.
func LoadEnv() {
	self.LoadEnv()
}

// LoadLocal loads mappings from the local configuration file.
func LoadLocal() error {
	return self.LoadLocal()
}

// LoadSystem loads mappings from the system configuration file.
func LoadSystem() error {
	return self.LoadSystem()
}

// Local gets the map of keys to values configured locally.
func Local() map[string]interface{} {
	return self.Local()
}

// Parse parses flags on the command line.
func Parse(arguments []string) error {
	return self.Parse(arguments)
}

// Prefix gets the environment variable prefix.
func Prefix() string {
	return self.Prefix()
}

// QueueStore queues a single key-value pair for storage into the local configuration file upon the next call to Save.
func QueueStore(key string, value interface{}) Configger {
	return self.QueueStore(key, value)
}

// Save atomically saves changes to the local configuration file.
func Save() error {
	return self.Save()
}

// SetAutoBind sets whether or not to use automatic environment variable binding.
func SetAutoBind(value bool) Configger {
	return self.SetAutoBind(value)
}

// SetPrefix sets the environment variable prefix.
func SetPrefix(value string) Configger {
	return self.SetPrefix(value)
}

// System gets the map of keys to values configured system-wide.
func System() map[string]interface{} {
	return self.System()
}

var self Configger
var mutex sync.Mutex

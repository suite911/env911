package config

import (
	"sync"

	"github.com/amy911/amy911/onfail"

	"github.com/amy911/env911/app"
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
func FlagSet() Flagger {
	return self.FlagSet()
}

// Get gets the value associated with key from the configuration sources
func Get(key string) interface{} {
	return self.Get(key)
}

// Load loads mappings from the several sources.
func Load() {
	self.Load()
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

// Prefix gets the environment variable prefix.
func Prefix() string {
	return self.Prefix()
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

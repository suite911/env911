package config

import (
	"sync"

	"github.com/amy911/env911/app"

	"github.com/amy911/error911/onfail"
	"github.com/amy911/flag911/flagger"
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

// Forward the call to the Flagger.
func Bool(name string, value bool, usage string) *bool {
	return self.Bool(name, value, usage)
}

// Forward the call to the Flagger.
func BoolP(name, shorthand string, value bool, usage string) *bool {
	return self.BoolP(name, shorthand, value, usage)
}

// Forward the call to the Flagger.
func BoolVar(p *bool, name string, value bool, usage string) {
	self.BoolVar(p, name, value, usage)
}

// Forward the call to the Flagger.
func BoolVarP(p *bool, name, shorthand string, value bool, usage string) {
	self.BoolVarP(p, name, shorthand, value, usage)
}

// Forward the call to the Flagger.
func Count(name string, by int, usage string) *int {
	return self.Count(name, by, usage)
}

// Forward the call to the Flagger.
func CountP(name, shorthand string, by int, usage string) *int {
	return self.CountP(name, shorthand, by, usage)
}

// Forward the call to the Flagger.
func CountVar(p *int, name string, by int, usage string) {
	self.CountVar(p, name, by, usage)
}

// Forward the call to the Flagger.
func CountVarP(p *int, name, shorthand string, by int, usage string) {
	self.CountVarP(p, name, shorthand, by, usage)
}

// Env gets the map of keys to values configured via environment variables.
func Env() map[string]interface{} {
	return self.Env()
}

// FlagSet gets the flag set used for configuration.
func FlagSet() flagger.Flagger {
	return self.FlagSet()
}

// Get gets the value associated with key from the configuration sources
func Get(key string) interface{} {
	return self.Get(key)
}

// Forward the call to the Flagger.
func Int(name string, value int, usage string) *int {
	return self.Int(name, value, usage)
}

// Forward the call to the Flagger.
func IntP(name, shorthand string, value int, usage string) *int {
	return self.IntP(name, shorthand, value, usage)
}

// Forward the call to the Flagger.
func IntVar(p *int, name string, value int, usage string) {
	self.IntVar(p, name, value, usage)
}

// Forward the call to the Flagger.
func IntVarP(p *int, name, shorthand string, value int, usage string) {
	self.IntVarP(p, name, shorthand, value, usage)
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

// Forward the call to the Flagger.
func String(name, value, usage string) *string {
	return self.String(name, value, usage)
}

// Forward the call to the Flagger.
func StringP(name, shorthand, value string, usage string) *string {
	return self.StringP(name, shorthand, value, usage)
}

// Forward the call to the Flagger.
func StringVar(p *string, name, value, usage string) {
	self.StringVar(p, name, value, usage)
}

// Forward the call to the Flagger.
func StringVarP(p *string, name, shorthand, value string, usage string) {
	self.StringVarP(p, name, shorthand, value, usage)
}

// System gets the map of keys to values configured system-wide.
func System() map[string]interface{} {
	return self.System()
}

var self Configger
var mutex sync.Mutex

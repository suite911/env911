package config

import (
	"github.com/amy911/env911/app"

	"gopkg.in/yaml.v2"
)

// Config describes the app configuration.
type Config struct {
	app     app.Apper
	flagSet Flagger

	env    map[string]interface{}
	local  map[string]interface{}
	system map[string]interface{}

	keys   [][2]string
	prefix string

	autoBind bool
}

// New creates a new Config.
func New(app app.Apper, flagSet Flagger, envPrefix string, args ...interface{}) *Config {
	return new(Config).Init(app, flagSet, envPrefix, args...)
}

// Init initializes a Config.
func (config *Config) Init(app app.Apper, flagSet Flagger, envPrefix string, args ...interface{}) *Config {
	config.app = app
	config.flagSet = flagSet
	config.envPrefix = envPrefix
	config.env = make(map[string]interface{})
	config.local = make(map[string]interface{})
	config.system = make(map[string]interface{})
	return config
}

// App gets the app being configured.
func (config *Config) App() app.Apper {
	return config.app
}

// AutoBind gets whether or not to use automatic environment variable binding.
func (config *Config) AutoBind() bool {
	return config.autoBind
}

// Bind binds the an environment key.
func (config *Config) Bind(key string) *Config {
	config.keys = append(config.keys, [2]string{config.prefix, key})
	return config
}

// Env gets the map of keys to values configured via environment variables.
func (config *Config) Env() map[string]interface{} {
	return config.env
}

// FlagSet gets the flag set used for configuration.
func (config *Config) FlagSet() Flagger {
	return config.flagSet
}

// Get gets the value associated with key from the configuration sources
func (config *Config) Get(key string) interface{} {
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
func (config *Config) Load() {
	config.LoadSystem()
	config.LoadLocal()
	config.LoadEnv()
}

// LoadEnv loads mappings from environment variables.
func (config *Config) LoadEnv() {
	for _, key := range config.keys {
		k := strings.ToUpper(key[0] + key[1])
		if v := os.Getenv(k); len(v) > 0 {
			config.env[key[1]] = v
		}
	}
}

// LoadLocal loads mappings from the local configuration file.
func (config *Config) LoadLocal() error {
	f, err := os.Open(config.App().LocalConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(config.local)
}

// LoadSystem loads mappings from the system configuration file.
func (config *Config) LoadSystem() error {
	f, err := os.Open(config.App().SystemConfigFile)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(config.system)
}

// Local gets the map of keys to values configured locally.
func (config *Config) Local() map[string]interface{} {
	return config.local
}

// Parse parses flags on the command line.
func (config *Config) Parse(arguments []string) error {
	return config.FlagSet().Parse(arguments)
}

// Prefix gets the environment variable prefix.
func (config *Config) Prefix() string {
	return config.prefix
}

// SetAutoBind sets whether or not to use automatic environment variable binding.
func (config *Config) SetAutoBind(value bool) *Config {
	config.autoBind = value
	return config
}

// SetPrefix sets the environment variable prefix.
func (config *Config) SetPrefix(value string) *Config {
	config.prefix = value
	return config
}

// System gets the map of keys to values configured system-wide.
func (config *Config) System() map[string]interface{} {
	return config.system
}

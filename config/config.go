package config

import (
	"github.com/amy911/env911/app"

	"gopkg.in/yaml.v2"
)

// Config describes the app configuration.
type Config struct {
	app     app.Apper
	flagSet Flagger

	envPrefix string
	keys      []string

	env    map[string]interface{}
	local  map[string]interface{}
	system map[string]interface{}
}

// New creates a new Config.
func New(app app.Apper, flagSet Flagger, envPrefix string, args ...interface{}) *Config {
	return new(Config).Init(app, flagSet, envPrefix, args...)
}

// AddEnv adds one or more environment keys from which to read.
func (config *Config) AddEnv(keys ...string) {
	config.keys = append(config.keys, keys)
}

// App gets the app being configured.
func (config *Config) App() app.Apper {
	return config.app
}

// Env gets the environment keys to read being configured.
func (config *Config) Env() []string {
	var keys []string
	for _, k := range config.keys {
		keys = append(keys, config.envPrefix+strings.ToUpper(k))
	}
	return keys
}

// EnvRaw gets the unmodified environment keys to read being configured.
func (config *Config) EnvRaw() []string {
	var keys []string
	for _, k := range config.keys {
		keys = append(keys, k)
	}
	return keys
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

// LoadConfig loads mappings from the configuration file.
func (config *Config) LoadConfigFile(m map[string]interface{}, path string, onFail ...onfail.OnFail) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(m)
}

// LoadEnv loads mappings from the environment.
func (config *Config) LoadEnv(m map[string]interface{}, onFail ...onfail.OnFail) {
	for _, k := range config.keys {
		if v := os.Getenv(config.envPrefix+strings.ToUpper(k)); len(v) > 0 {
			config.map_[k] = v
		}
	}
}

// Env gets the map of keys to values configured via environment variables.
func (config *Config) Env() map[string]interface{} {
	return config.env
}

// Local gets the map of keys to values configured locally
func (config *Config) Local() map[string]interface{} {
	return config.local
}

// System gets the map of keys to values configured system-wide
func (config *Config) System() map[string]interface{} {
	return config.system
}

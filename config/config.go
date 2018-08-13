package config

import (
	"github.com/amy911/env911/app"

	"gopkg.in/yaml.v2"
)

// Config describes the app configuration.
type Config struct {
	app       app.Apper
	env       []string
	envPrefix string
	map_      map[string]interface{}
}

// New creates a new Config.
func New(app app.Apper, envPrefix string, args ...interface{}) *Config {
	return new(Config).Init(app, envPrefix, args...)
}

// AddEnv adds one or more environment keys from which to read.
func (config *Config) AddEnv(keys ...string) {
	config.env = append(config.env, keys)
}

// App gets the app being configured.
func (config *Config) App() app.Apper {
	return config.app
}

// Env gets the environment keys to read being configured.
func (config *Config) Env() []string {
	var env []string
	for _, k := range config.env {
		env = append(env, config.envPrefix+strings.ToUpper(k))
	}
	return env
}

// EnvRaw gets the unmodified environment keys to read being configured.
func (config *Config) EnvRaw() []string {
	var env []string
	for _, k := range config.env {
		env = append(env, k)
	}
	return env
}

// Init initializes a Config.
func (config *Config) Init(app app.Apper, envPrefix string, args ...interface{}) *Config {
	config.app = app
	config.envPrefix = envPrefix
	config.map_ = make(map[string]interface{})
	return config
}

// LoadConfig loads mappings from the configuration file.
func (config *Config) LoadConfig() error {
	f, err := os.Open(config.App().ConfigFile())
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(config.map_)
}

// LoadEnv loads mappings from the environment.
func (config *Config) LoadEnv() {
	for _, k := range config.env {
		if v := os.Getenv(config.envPrefix+strings.ToUpper(k)); len(v) > 0 {
			config.map_[k] = v
		}
	}
}

// Map gets the map of keys to configured values.
func (config *Config) Map() map[string]interface{} {
	return config.map_
}

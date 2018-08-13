package config

import (
	"github.com/amy911/env911/app"

	"gopkg.in/yaml.v2"
)

// Config describes the app configuration.
type Config struct {
	app       app.Apper
	envKeys   []string
	envPrefix string

	env    map[string]interface{}
	local  map[string]interface{}
	system map[string]interface{}
}

// New creates a new Config.
func New(app app.Apper, envPrefix string, args ...interface{}) *Config {
	return new(Config).Init(app, envPrefix, args...)
}

// AddEnv adds one or more environment keys from which to read.
func (config *Config) AddEnv(keys ...string) {
	config.envKeys = append(config.envKeys, keys)
}

// App gets the app being configured.
func (config *Config) App() app.Apper {
	return config.app
}

// Env gets the environment keys to read being configured.
func (config *Config) Env() []string {
	var envKeys []string
	for _, k := range config.envKeys {
		envKeys = append(envKeys, config.envPrefix+strings.ToUpper(k))
	}
	return envKeys
}

// EnvRaw gets the unmodified environment keys to read being configured.
func (config *Config) EnvRaw() []string {
	var envKeys []string
	for _, k := range config.envKeys {
		envKeys = append(envKeys, k)
	}
	return envKeys
}

// Init initializes a Config.
func (config *Config) Init(app app.Apper, envPrefix string, args ...interface{}) *Config {
	config.app = app
	config.envPrefix = envPrefix
	config.map_ = make(map[string]interface{})
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
	for _, k := range config.envKeys {
		if v := os.Getenv(config.envPrefix+strings.ToUpper(k)); len(v) > 0 {
			config.map_[k] = v
		}
	}
}

// Map gets the map of keys to configured values.
func (config *Config) Map() map[string]interface{} {
	return config.map_
}

package config

import (
	"os"
	"strings"

	"github.com/amy911/env911/app"
	"github.com/amy911/env911/safesave"

	"github.com/amy911/flag911/flag"

	"github.com/ogier/pflag"
	"gopkg.in/yaml.v2"
)

// Config describes the app configuration.
type Config struct {
	app     app.Apper
	flagSet *flag.FlagSet

	env    map[string]interface{}
	local  map[string]interface{}
	system map[string]interface{}

	keys   [][2]string
	prefix string

	autoBind bool
}

// New creates a new Config.
func New(prefix string, flagSet *flag.FlagSet, app app.Apper) *Config {
	return new(Config).Init(prefix, flagSet, app)
}

// Init initializes a Config.
func (config *Config) Init(prefix string, flagSet *flag.FlagSet, app app.Apper) *Config {
	if flagSet == nil {
		flagSet = flag.New(pflag.CommandLine)
	}
	if app == nil {
		panic("You must specify an app")
	}
	config.app = app
	config.flagSet = flagSet
	config.prefix = prefix
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
func (config *Config) Bind(key string) Configger {
	config.keys = append(config.keys, [2]string{config.prefix, key})
	return config
}

// Env gets the map of keys to values configured via environment variables.
func (config *Config) Env() map[string]interface{} {
	return config.env
}

// FlagSet gets the flag set used for configuration.
func (config *Config) FlagSet() *flag.FlagSet {
	return config.flagSet
}

// Get gets the value associated with key from the configuration sources
func (config *Config) Get(key string) interface{} {
	if v, ok := config.Env()[key]; ok {
		return v
	}
	if v, ok := config.Local()[key]; ok {
		return v
	}
	if v, ok := config.System()[key]; ok {
		return v
	}
	return nil
}

// Load loads mappings from the several sources.
func (config *Config) Load() Configger {
	config.LoadSystem()
	config.LoadLocal()
	config.LoadEnv()
	return config
}

// LoadAndParse loads mappings from the several sources and parses flags on the command line.
func (config *Config) LoadAndParse() {
	config.Load()
	config.Parse(os.Args)
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
	f, err := os.Open(config.App().LocalConfigFile())
	if err != nil {
		return err
	}
	defer f.Close()
	return yaml.NewDecoder(f).Decode(config.local)
}

// LoadSystem loads mappings from the system configuration file.
func (config *Config) LoadSystem() error {
	f, err := os.Open(config.App().SystemConfigFile())
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

// QueueStore queues a single key-value pair for storage into the local configuration file upon the next call to Save.
func (config *Config) QueueStore(key string, value interface{}) Configger {
	config.Local()[key] = value
	return config
}

// Save atomically saves changes to the local configuration file.
func (config *Config) Save() error {
	ss, err := safesave.New(config.App().LocalConfigFile(), config.App().LocalConfigNew())
	if err != nil {
		return err
	}
	enc := yaml.NewEncoder(ss)
	if err := enc.Encode(config.local); err != nil {
		enc.Close()
		ss.Close()
		return err
	}
	if err := enc.Close(); err != nil {
		ss.Close()
		return err
	}
	return ss.Close()
}

// SetAutoBind sets whether or not to use automatic environment variable binding.
func (config *Config) SetAutoBind(value bool) Configger {
	config.autoBind = value
	if value {
		config.FlagSet().SetHook("github.com/amy911/env911/config", flagHookAutoBind, config)
	} else {
		config.FlagSet().DeleteHook("github.com/amy911/env911/config")
	}
	return config
}

// SetPrefix sets the environment variable prefix.
func (config *Config) SetPrefix(value string) Configger {
	config.prefix = value
	return config
}

// System gets the map of keys to values configured system-wide.
func (config *Config) System() map[string]interface{} {
	return config.system
}

func flagHookAutoBind(
	fs *FlagSet,
	fn string,
	p interface{},
	name string,
	shorthand string,
	value interface{},
	usage string,
	user []interface{},
) error {
	if len(user) != 1 {
		panic("Issue with hijacked hook")
	}
	config, ok := user[0].(*Config)
	if !ok {
		panic("Issue with hijacked hook")
	}
	if !config.AutoBind() {
		panic("Issue with SetAutoBind or flagHookAutoBind")
	}
	config.Bind(name)
	return nil
}

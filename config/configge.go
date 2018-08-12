package config

// Configger is the interface for app configuration.
type Configger interface {

	// App gets the app being configured.
	App() app.Apper

	// Map gets the map of keys to configured values.
	Map() map[string]interface{}

}

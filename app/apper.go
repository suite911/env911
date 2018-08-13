package app

// Apper is the interface for an app.
type Apper interface {

	// Cache gets the full path to the local app cache directory.
	Cache() string

	// Config gets the full path to the local app configuration directory.
	Config() string

	// ConfigFile gets the full path to the local app configuration file.
	ConfigFile() string

	// Data gets the full path to the local app data directory.
	Data() string

	// Desktop gets the full path to the current user's Desktop directory.
	Desktop() string

	// Documents gets the full path to the current user's Documents directory.
	Documents() string

	// Downloads gets the full path to the current user's Downloads directory.
	Downloads() string

	// Home gets the full path to the current user's Home directory.
	Home() string

	// Name gets the app name.
	Name() string

	// Path gets the app path.
	Path() string

	// Pictures gets the full path to the current user's Pictures directory.
	Pictures() string

	// Screenshots gets the full path to the current user's Screenshots directory.
	Screenshots() string

	// Vendor gets the app vendor.
	Vendor() string

}

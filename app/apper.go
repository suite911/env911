package app

// Apper is the interface for an app.
type Apper interface {

	// Cache gets the full path to the local app cache directory.
	Cache() string

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

	// LocalConfig gets the full path to the local app configuration directory.
	LocalConfig() string

	// LocalConfigFile gets the full path to the local app configuration file.
	LocalConfigFile() string

	// Path gets the app path.
	Path() string

	// Pictures gets the full path to the current user's Pictures directory.
	Pictures() string

	// Screenshots gets the full path to the current user's Screenshots directory.
	Screenshots() string

	// SystemConfig gets the full path to the system app configuration directory.
	SystemConfig() string

	// SystemConfigFile gets the full path to the system app configuration file.
	SystemConfigFile() string

	// Vendor gets the app vendor.
	Vendor() string

}

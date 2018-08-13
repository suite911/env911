package app

// Apper is the interface for an app.
type Apper interface {

	// Bin returns the full path to the directory in which the running executable is located.
	Bin() string

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

	// Exe returns the full path to the running executable after attempting to expand symlinks.
	Exe() string

	// Home gets the full path to the current user's Home directory.
	Home() string

	// Man gets the full path to the man directory.
	Man() string

	// Name gets the app name.
	Name() string

	// LocalConfig gets the full path to the local app configuration directory.
	LocalConfig() string

	// LocalConfigFile gets the full path to the local app configuration file.
	LocalConfigFile() string

	// LocalConfigNew returns the full path to a write in-progress copy of the local app configuration file.
	LocalConfigNew() string

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

package app

import "path/filepath"

// App is the type of an app.
type App struct {
	vendor, name, path string

	bin              string
	cache            string
	data             string
	desktop          string
	documents        string
	downloads        string
	home             string
	localConfig      string
	localConfigFile  string
	man              string
	pictures         string
	screenshots      string
	systemConfig     string
	systemConfigFile string
}

// New creates a new App.
func New(args ...interface{}) *App {
	return new(App).Init(args...)
}

// Init initializes an App.
func (app *App) Init(args ...interface{}) *App {
	var pathElems []string
	for _, arg := range args {
		if str, ok := arg.(string); ok {
			if len(str) < 1 {
				continue
			}
			pathElems = append(pathElems, str)
			if len(app.vendor) < 1 {
				app.vendor = app.name
			}
			app.name = str
		}
	}
	app.path = filepath.Join(pathElems)
	exeWithSymlinks, err := os.Executable()
	app.bin = filepath.Dir(exeWithSymlinks)
	exeMaybeWithoutSymlinks := exeWithSymlinks
	if err == nil {
		if p, err = filepath.EvalSymlinks(exeWithSymlinks); err == nil {
			exeMaybeWithoutSymlinks = p
		}
	}
	app.exe = exeMaybeWithoutSymlinks
	app.osInit(args...)
	app.localConfigFile = filepath.Join(app.localConfig, "config.yml")
	app.systemConfigFile = filepath.Join(app.systemConfig, "config.yml")
	return app
}

// Bin returns the full path to the directory in which the running executable is located.
func (app App) Bin() string {
	return app.bin
}

// Cache returns the full path to a local app cache directory.
func (app App) Cache() string {
	return app.cache
}

// Data returns the full path to a local app data directory.
func (app App) Data() string {
	return app.data
}

// Desktop returns the full path to the desktop directory.
func (app App) Desktop() string {
	return app.desktop
}

// Documents returns the full path to the documents directory.
func (app App) Documents() string {
	return app.documents
}

// Downloads returns the full path to the downloads directory.
func (app App) Downloads() string {
	return app.downloads
}

// Home returns the full path to the home directory.
func (app App) Home() string {
	return app.home
}

// Man gets the full path to the man directory.
func (app App) Man() string {
	return app.man
}

// LocalConfig returns the full path to a local app configuration directory.
func (app App) LocalConfig() string {
	return app.localConfig
}

// LocalConfigFile returns the full path to a local app configuration file.
func (app App) LocalConfigFile() string {
	return app.localConfigFile
}

// Name returns the app name.
func (app App) Name() string {
	return app.name
}

// Path returns the app path.
func (app App) Path() string {
	return app.path
}

// Pictures returns the full path to the pictures directory.
func (app App) Pictures() string {
	return app.pictures
}

// Screenshots returns the full path to the screenshots directory.
func (app App) Screenshots() string {
	return app.screenshots
}

// SystemConfig returns the full path to a system app configuration directory.
func (app App) SystemConfig() string {
	return app.systemConfig
}

// SystemConfigFile returns the full path to a system app configuration file.
func (app App) SystemConfigFile() string {
	return app.systemConfigFile
}

// Vendor returns the app vendor.
func (app App) Vendor() string {
	return app.vendor
}

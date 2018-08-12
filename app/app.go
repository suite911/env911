package app

import "path/filepath"

// App is the type of an app.
type App struct {
	vendor, name, path string

	cache       string
	config      string
	data        string
	desktop     string
	documents   string
	downloads   string
	home        string
	pictures    string
	screenshots string
}

// New creates a new App.
func New(args ...interface{}) *App {
	return new(App).Init(args...)
}

// Cache returns the full path to a local app cache directory.
func (app App) Cache() string {
	return app.cache
}

// Config returns the full path to a local app config directory.
func (app App) Config() string {
	return app.config
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
	app.osInit(args...)
	return app
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

// Vendor returns the app vendor.
func (app App) Vendor() string {
	return app.vendor
}

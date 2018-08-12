package app

import "sync"

// Self is this app.
var Self Apper

// Cache returns the full path to a local app cache directory.
func Cache() string {
	return app.Self.Cache()
}

// Config returns the full path to a local app config directory.
func Config() string {
	return app.Self.Config()
}

// Data returns the full path to a local app data directory.
func Data() string {
	return app.Self.Data()
}

// Desktop returns the full path to the desktop directory.
func Desktop() string {
	return app.Self.Desktop()
}

// Documents returns the full path to the documents directory.
func Documents() string {
	return app.Self.Documents()
}

// Downloads returns the full path to the downloads directory.
func Downloads() string {
	return app.Self.Downloads()
}

// Home returns the full path to the home directory.
func Home() string {
	return app.Self.Home()
}

// Init initializes this app.
func Init(args ...interface{}) {
	mutexPath.Lock(); defer mutexPath.Unlock()
	if Self != nil {
		return
	}
	Self = New(args...)
}

// Name returns the app name.
func Name() string {
	return app.Self.Name()
}

// Path returns the app path.
func Path() string {
	return app.Self.Path()
}

// Pictures returns the full path to the pictures directory.
func Pictures() string {
	return app.Self.Pictures()
}

// Screenshots returns the full path to the screenshots directory.
func Screenshots() string {
	return app.Self.Screenshots()
}

// Vendor returns the app vendor.
func Vendor() string {
	return app.Self.Vendor()
}

var mutexPath sync.Mutex

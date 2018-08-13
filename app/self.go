package app

import (
	"sync"

	"github.com/amy911/amy911/onfail"
)

// Init initializes this app.
func Init(app Apper, onFail ...onfail.OnFail) {
	mutex.Lock(); defer mutex.Unlock()
	if self != nil {
		onfail.Fail("Double init", self, onfail.Panic, onFail)
		return
	}
	self = app
}

// Cache returns the full path to a local app cache directory.
func Cache() string {
	return self.Cache()
}

// Config returns the full path to a local app config directory.
func Config() string {
	return self.Config()
}

// Data returns the full path to a local app data directory.
func Data() string {
	return self.Data()
}

// Desktop returns the full path to the desktop directory.
func Desktop() string {
	return self.Desktop()
}

// Documents returns the full path to the documents directory.
func Documents() string {
	return self.Documents()
}

// Downloads returns the full path to the downloads directory.
func Downloads() string {
	return self.Downloads()
}

// Home returns the full path to the home directory.
func Home() string {
	return self.Home()
}

// Man returns the full path to the man directory.
func Man() string {
	return self.Man()
}

// Name returns the app name.
func Name() string {
	return self.Name()
}

// Path returns the app path.
func Path() string {
	return self.Path()
}

// Pictures returns the full path to the pictures directory.
func Pictures() string {
	return self.Pictures()
}

// Screenshots returns the full path to the screenshots directory.
func Screenshots() string {
	return self.Screenshots()
}

// Vendor returns the app vendor.
func Vendor() string {
	return self.Vendor()
}

var self Apper
var mutex sync.Mutex

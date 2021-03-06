package app

import (
	"sync"

	"github.com/suite911/error911/onfail"
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

// IsInit reports whether or not this app has been initialized.
func IsInit() bool {
	return self != nil
}

// Bin returns the full path to the directory in which the running executable is located.
func Bin() string {
	return self.Bin()
}

// Cache returns the full path to the local app cache directory.
func Cache() string {
	return self.Cache()
}

// Data returns the full path to the local app data directory.
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

// Exe returns the full path to the running executable after attempting to expand symlinks.
func Exe() string {
	return self.Exe()
}

// Home returns the full path to the home directory.
func Home() string {
	return self.Home()
}

// LocalConfig returns the full path to a local app configuration directory.
func LocalConfig() string {
	return self.LocalConfig()
}

// LocalConfigFile returns the full path to a local app configuration directory.
func LocalConfigFile() string {
	return self.LocalConfigFile()
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

// SystemConfig returns the full path to the system app configuration directory.
func SystemConfig() string {
	return self.SystemConfig()
}

// SystemConfigFile returns the full path to the system app configuration directory.
func SystemConfigFile() string {
	return self.SystemConfigFile()
}

// Vendor returns the app vendor.
func Vendor() string {
	return self.Vendor()
}

var self Apper
var mutex sync.Mutex

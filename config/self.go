package config

import (
	"sync"

	"github.com/amy911/env911/config/app"
)

// Self is the configuration for this app.
var Self Configger

// Init initializes the configuration for this app.
func Init(app app.Apper, args ...interface{}) {
	mutexPath.Lock(); defer mutexPath.Unlock()
	if Self != nil {
		return
	}
	Self = New(app, args...)
}

var mutexPath sync.Mutex

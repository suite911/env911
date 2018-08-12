package app

import "sync"

var Self *App

func Init(args ...interface{}) {
	mutexPath.Lock(); defer mutexPath.Unlock()
	if Self != nil {
		return
	}
	Self = New(args...)
}

var mutexPath sync.Mutex

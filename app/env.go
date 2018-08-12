package app

import "sync"

var Env *AppEnv

func Init(args ...interface{}) {
	mutexPath.Lock(); defer mutexPath.Unlock()
	if Env != nil {
		return
	}
	Env = New(args...)
}

var mutexPath sync.Mutex

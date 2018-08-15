package env911

import (
	"github.com/amy911/env911/app"
	"github.com/amy911/env911/config"
)

// InitAll initializes every part of env911 at once.
func InitAll(prefix string, flagSet config.Flagger, path ...string) {
	a := app.New(path...)
	app.Init(a)
	c := config.New(prefix, flagSet, a)
	config.Init(c)
}

// IsInitAll reports whether or not every part of env911 has been initialized.
func IsInitAll() bool {
	return app.IsInit() && config.IsInit()
}

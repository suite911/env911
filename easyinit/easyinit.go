package easyinit

import (
	"github.com/amy911/env911/app"
	"github.com/amy911/env911/config"
)

func Init(prefix string, flagSet config.Flagger, path ...string) {
	a := app.New(path...)
	app.Init(a)
	c := config.New(prefix, flagSet, a)
	config.Init(c)
}

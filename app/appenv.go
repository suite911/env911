package app

import "path/filepath"

type Type struct {
	vendor, name, path string
}

func New(args ...interface{}) *Type {
	return new(Type).Init(args...)
}

func (app *Type) Init(args ...interface{}) *Type {
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

func (app *Type) Name() string {
	return app.name
}

func (app *Type) Path() string {
	return app.path
}

func (app *Type) Vendor() string {
	return app.vendor
}

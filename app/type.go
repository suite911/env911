package app

import "path/filepath"

type Type struct {
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

func New(args ...interface{}) *Type {
	return new(Type).Init(args...)
}

func (app *Type) Cache() string {
	return app.cache
}

func (app *Type) Config() string {
	return app.config
}

func (app *Type) Data() string {
	return app.data
}

func (app *Type) Desktop() string {
	return app.desktop
}

func (app *Type) Documents() string {
	return app.documents
}

func (app *Type) Downloads() string {
	return app.downloads
}

func (app *Type) Home() string {
	return app.home
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

func (app *Type) Pictures() string {
	return app.pictures
}

func (app *Type) Screenshots() string {
	return app.screenshots
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

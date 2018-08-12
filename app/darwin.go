// +build darwin

package app

import (
	"os"
	"path/filepath"
)

func (app *App) osInit(args ...interface{}) {
	home := os.Getenv("HOME")
	if len(home) < 1 {
		home = "/root"
	}
	xdg_cache_home := os.Getenv("XDG_CACHE_HOME")
	if len(xdg_cache_home) < 1 {
		xdg_cache_home = filepath.Join(home, ".cache")
	}
	xdg_config_home := os.Getenv("XDG_CONFIG_HOME")
	if len(xdg_config_home) < 1 {
		xdg_config_home = filepath.Join(home, ".config")
	}
	xdg_data_home := os.Getenv("XDG_DATA_HOME")
	if len(xdg_data_home) < 1 {
		xdg_data_home = filepath.Join(home, ".local/share")
	}

	app.cache = filepath.Join(xdg_cache_home, app.path)
	app.config = filepath.Join(xdg_config_home, app.path)
	app.data = filepath.Join(xdg_data_home, app.path)
	app.desktop = filepath.Join(home, "Desktop")
	app.documents = filepath.Join(home, "Documents")
	app.downloads = filepath.Join(home, "Downloads")
	app.home = home
	app.pictures = filepath.Join(home, "Pictures")
	app.screenshots = filepath.Join(app.pictures, "Screenshots")
}

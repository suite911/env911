package app

import (
	"os"
	"path/filepath"
)

func SystemPath(bin, share string) string {
	local := filepath.Dir(bin)
	if len(local) == 1 && os.IsPathSeparator(local[0]) {
		local = filepath.Join(local, "usr")
	}
	return filepath.Join(local, share)
}

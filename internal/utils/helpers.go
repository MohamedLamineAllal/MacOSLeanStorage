package utils

import (
	"os"
	"path/filepath"
)

// GetAppCacheDir returns the persistent cache directory for the application.
// It falls back to the system temp directory if the user cache dir is unavailable.
func GetAppCacheDir() string {
	cache, err := os.UserCacheDir()
	if err != nil {
		cache = os.TempDir()
	}
	path := filepath.Join(cache, "mls")
	_ = os.MkdirAll(path, 0755)
	return path
}

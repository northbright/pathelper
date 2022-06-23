package pathelper

import (
	"os"
	"path/filepath"
)

// ExecDir returns the absolute path of given relative path to the current executable dir.
// If relPath is "", it just returns the absolute path of current executable dir.
func ExecDir(relPath string) (string, error) {
	// os.Executable requires Go 1.18+
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Join(filepath.Dir(ex), relPath), nil
}

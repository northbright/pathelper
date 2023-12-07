// pathelper is a golang package which provides path related helper functions.
package pathelper

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
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

// FileExists returns if given file exists or not.
func FileExists(f string) bool {
	_, err := os.Stat(f)
	return !errors.Is(err, fs.ErrNotExist)
}

// CommandExists returns if given command exists or not.
func CommandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// CreateDirIfNotExists creates the directory if it does not exists.
func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	if !FileExists(dir) {
		return os.MkdirAll(dir, perm)
	}
	return nil
}

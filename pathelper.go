// pathelper is a golang package which provides path related helper functions.
package pathelper

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

// ExecDir returns the absolute path of current executable dir.
// For Go 1.18 and later, you may just use os.Executable instead.
func ExecDir() (string, error) {
	// os.Executable requires Go 1.18+
	return os.Executable()
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

// BaseWithoutExt returns file base name without ext name.
func BaseWithoutExt(fileName string) string {
	base := filepath.Base(fileName)
	ext := filepath.Ext(fileName)

	return base[:len(base)-len(ext)]
}

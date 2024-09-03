// pathelper is a golang package which provides path related helper functions.
package pathelper

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ExecDir returns the absolute path of current executable's dir.
func ExecDir() (string, error) {
	// os.Executable requires Go 1.18+
	p, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(p), nil
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

// ReplacePrefix returns a copy of path that old prefix in the path replaced with new prefix.
// The returned path is converted into an operating system path.
// In the function, it creates slash-separated copies of path and old / prefix then do the replacement.
// So path, oldPrefix and newPrefix can be slash-separated or operating system paths.
// e.g. path = "..\..\a\b\c", prefix = "../../a", new prefix = "C:\Users\RUNNER~1\AppData\Local\a"
// It returns "C:\Users\RUNNER~1\AppData\Local\a\b\c" on Windows.
func ReplacePrefix(path, oldPrefix, newPrefix string) string {
	path = filepath.ToSlash(filepath.Clean(path))
	oldPrefix = filepath.ToSlash(filepath.Clean(oldPrefix))
	newPrefix = filepath.ToSlash(filepath.Clean(newPrefix))

	path = strings.Replace(path, oldPrefix, newPrefix, 1)
	return filepath.FromSlash(path)
}

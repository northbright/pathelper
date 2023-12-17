package pathelper_test

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/northbright/pathelper"
)

// Run "go test -c && ./pathelper.test"
func ExampleExecDir() {
	// Get current executable dir.
	execDir, _ := pathelper.ExecDir()
	log.Printf("executable dir: %v", execDir)

	// Make an absolute path for image dir.
	imageDir := filepath.Join(execDir, "static", "images")
	log.Printf("image dir: %v", imageDir)

	// Output:
}

func ExampleFileExists() {
	exists := pathelper.FileExists("/tmp")
	log.Printf("/tmp exists: %v", exists)

	// Output:
}

func ExampleCommandExists() {
	exists := pathelper.CommandExists("go")
	log.Printf("go exists: %v", exists)

	// Output:
}

func ExampleCreateDirIfNotExists() {
	dir := "test/output"
	if err := pathelper.CreateDirIfNotExists(dir, 0755); err != nil {
		log.Printf("CreateDirIfNotExists error: %v", err)
		return
	}

	exists := pathelper.FileExists(dir)
	fmt.Printf("%v\n", exists)

	// Output:
	// true
}

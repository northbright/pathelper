package pathelper_test

import (
	"log"

	"github.com/northbright/pathelper"
)

// Run "go test -c && ./pathelper.test"
func ExampleExecDir() {
	// Pass "" as relPath to get current executable dir.
	execDir, _ := pathelper.ExecDir("")
	log.Printf("exectuable dir: %v", execDir)

	// Pass "./img" as relPath to get absolute path of image dir.
	imgDir, _ := pathelper.ExecDir("./img")
	log.Printf("image dir: %v", imgDir)

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

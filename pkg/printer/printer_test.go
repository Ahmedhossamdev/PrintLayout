package printer

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestPrintProjectStructure tests the PrintProjectStructure function.
func TestPrintProjectStructure(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create the folder structure dynamically
	createTestProjectStructure(t, tmpDir)

	// Change the working directory to the temporary directory
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir) // Restore the original working directory
	os.Chdir(tmpDir)

	// Capture the output
	output := captureOutput(func() {
		PrintProjectStructure()
	})

	// Get the base name of the temporary directory
	rootName := filepath.Base(tmpDir)

	// Define the expected output based on the created structure
	expected := rootName + "/\n" +
		"├── cmd/\n" +
		"│   └── main.go\n" +
		"├── go.mod\n" +
		"├── internal/\n" +
		"│   └── utils/\n" +
		"│       └── utils.go\n" +
		"└── pkg/\n" +
		"    └── printer/\n" +
		"        ├── printer.go\n" +
		"        └── printer_test.go\n"
	// Normalize the output and expected strings
	output = strings.TrimSpace(output)
	expected = strings.TrimSpace(expected)

	// Compare the output
	if output != expected {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
	}
}

// createTestProjectStructure creates a sample project structure for testing.
// createTestProjectStructure creates a sample project structure for testing.
func createTestProjectStructure(t *testing.T, root string) {
	// Define the directories to create
	dirs := []string{
		"cmd",
		"internal/utils",
		"pkg/printer",
	}

	// Define the files to create
	files := []string{
		"cmd/main.go",
		"internal/utils/utils.go",
		"pkg/printer/printer.go",
		"pkg/printer/printer_test.go",
		"go.mod",
	}

	// Create directories
	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(root, dir), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
	}

	// Create files
	for _, file := range files {
		f, err := os.Create(filepath.Join(root, file))
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
		f.Close()
	}
}

// captureOutput captures the output printed to stdout.
func captureOutput(f func()) string {
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	out, _ := io.ReadAll(r)
	os.Stdout = rescueStdout

	return string(out)
}

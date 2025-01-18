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
	tmpDir := t.TempDir()

	createTestProjectStructure(t, tmpDir)

	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir) // Restore the original working directory
	os.Chdir(tmpDir)

	output := captureOutput(func() {
		PrintProjectStructure(".", "")
	})

	rootName := filepath.Base(tmpDir)

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
	output = strings.TrimSpace(output)
	expected = strings.TrimSpace(expected)

	if output != expected {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
	}
}

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

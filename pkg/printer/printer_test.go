package printer

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrintProjectStructure(t *testing.T) {
	tmpDir := t.TempDir()

	createTestProjectStructure(t, tmpDir)

	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir)
	os.Chdir(tmpDir)

	output := captureOutput(func() {
		PrintProjectStructure()
	})

	rootName := filepath.Base(tmpDir)

	// TODO: Update the expected output (more smartly)
	expected := rootName + "/\n" +
		"├── main.go\n" +
		"├── pkg/\n" +
		"│   └── printer/\n" +
		"│       ├── printer_test.go\n" +
		"│       └── printer.go\n" +
		"├── cmd/\n" +
		"│   └── main.go\n" +
		"├── go.mod\n" +
		"└── internal/\n" +
		"    └── utils/\n" +
		"        └── utils.go\n"

	output = strings.TrimSpace(output)
	expected = strings.TrimSpace(expected)

	if output != expected {
		t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
	}
}

func createTestProjectStructure(t *testing.T, root string) {
	dirs := []string{
		"cmd",
		"internal/utils",
		"pkg/printer",
	}

	files := []string{
		"cmd/main.go",
		"internal/utils/utils.go",
		"pkg/printer/printer.go",
		"pkg/printer/printer_test.go",
		"go.mod",
		"main.go",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(root, dir), 0755)
		if err != nil {
			t.Fatalf("Failed to create directory: %v", err)
		}
	}

	for _, file := range files {
		f, err := os.Create(filepath.Join(root, file))
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}
		f.Close()
	}
}

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

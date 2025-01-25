package printer

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"gopkg.in/yaml.v3"
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

	// Test text output
	t.Run("TextOutput", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "name", "asc", false)
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
			"        └── printer_test.go\n" +
			"\n5 directories, 5 files\n"
		output = strings.TrimSpace(output)
		expected = strings.TrimSpace(expected)

		if output != expected {
			t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
		}
	})

	// Test JSON output
	t.Run("JSONOutput", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "json", "blue", "green", "red", []string{}, "name", "asc", false)
		})

		// Verify that the output is valid JSON
		var result interface{}
		if err := json.Unmarshal([]byte(output), &result); err != nil {
			t.Errorf("Output is not valid JSON: %v", err)
		}
	})

	// Test XML output
	t.Run("XMLOutput", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "xml", "blue", "green", "red", []string{}, "name", "asc", false)
		})

		// Verify that the output is valid XML
		var result interface{}
		if err := xml.Unmarshal([]byte(output), &result); err != nil {
			t.Errorf("Output is not valid XML: %v", err)
		}
	})

	// Test YAML output
	t.Run("YAMLOutput", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "yaml", "blue", "green", "red", []string{}, "name", "asc", false)
		})

		// Verify that the output is valid YAML
		var result interface{}
		if err := yaml.Unmarshal([]byte(output), &result); err != nil {
			t.Errorf("Output is not valid YAML: %v", err)
		}
	})

	// Test exclusion patterns
	t.Run("ExclusionPatterns", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{"*.go"}, "name", "asc", false)
		})

		rootName := filepath.Base(tmpDir)

		expected := rootName + "/\n" +
			"├── cmd/\n" +
			"├── go.mod\n" +
			"├── internal/\n" +
			"│   └── utils/\n" +
			"└── pkg/\n" +
			"    └── printer/\n" +
			"\n5 directories, 1 files\n"
		output = strings.TrimSpace(output)
		expected = strings.TrimSpace(expected)

		if output != expected {
			t.Errorf("Unexpected output:\nGot:\n%s\nExpected:\n%s", output, expected)
		}
	})

	// Test sorting by name (ascending)
	t.Run("SortByNameAsc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "name", "asc", false)
		})

		// Verify that the output is sorted by name in ascending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})

	// Test sorting by name (descending)
	t.Run("SortByNameDesc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "name", "desc", false)
		})

		// Verify that the output is sorted by name in descending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})

	// Test sorting by size (ascending)
	t.Run("SortBySizeAsc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "size", "asc", false)
		})

		// Verify that the output is sorted by size in ascending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})

	// Test sorting by size (descending)
	t.Run("SortBySizeDesc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "size", "desc", false)
		})

		// Verify that the output is sorted by size in descending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})

	// Test sorting by time (ascending)
	t.Run("SortByTimeAsc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "time", "asc", false)
		})

		// Verify that the output is sorted by time in ascending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})

	// Test sorting by time (descending)
	t.Run("SortByTimeDesc", func(t *testing.T) {
		output := captureOutput(func() {
			PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "time", "desc", false)
		})

		// Verify that the output is sorted by time in descending order
		// You can add specific checks based on your expected output
		t.Log(output)
	})
}

// createTestProjectStructure creates a sample project structure for testing.
func createTestProjectStructure(tb testing.TB, root string) {
	// Define the directories to create
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
	}

	for _, dir := range dirs {
		err := os.MkdirAll(filepath.Join(root, dir), 0755)
		if err != nil {
			tb.Fatalf("Failed to create directory: %v", err)
		}
	}

	for _, file := range files {
		f, err := os.Create(filepath.Join(root, file))
		if err != nil {
			tb.Fatalf("Failed to create file: %v", err)
		}
		f.Close()

		if strings.HasSuffix(file, ".go") {
			modTime := time.Now().Add(-time.Hour * 24) // Set to 24 hours ago
			os.Chtimes(filepath.Join(root, file), modTime, modTime)
		}
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

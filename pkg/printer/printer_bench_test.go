package printer

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

// BenchmarkPrintProjectStructure benchmarks the PrintProjectStructure function.
func BenchmarkPrintProjectStructure(b *testing.B) {
	// Create a temporary directory for benchmarking
	tmpDir := b.TempDir()
	createTestProjectStructure(b, tmpDir)

	// Change to the temporary directory
	oldDir, err := os.Getwd()
	if err != nil {
		b.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir) // Restore the original working directory
	os.Chdir(tmpDir)

	// Run the benchmark
	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "name", "asc")
	}
}

// BenchmarkPrintProjectStructure_JSON benchmarks the JSON output format.
func BenchmarkPrintProjectStructure_JSON(b *testing.B) {
	// Create a temporary directory for benchmarking
	tmpDir := b.TempDir()
	createTestProjectStructure(b, tmpDir)

	oldDir, err := os.Getwd()
	if err != nil {
		b.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir) // Restore the original working directory
	os.Chdir(tmpDir)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		PrintProjectStructure(".", "", "", false, "json", "blue", "green", "red", []string{}, "name", "asc")
	}
}

// BenchmarkPrintProjectStructure_LargeDirectory benchmarks performance with a large directory.
func BenchmarkPrintProjectStructure_LargeDirectory(b *testing.B) {
	tmpDir := b.TempDir()
	createLargeTestProjectStructure(b, tmpDir)

	oldDir, err := os.Getwd()
	if err != nil {
		b.Fatalf("Failed to get current working directory: %v", err)
	}
	defer os.Chdir(oldDir) // Restore the original working directory
	os.Chdir(tmpDir)

	b.ResetTimer() // Reset the timer to exclude setup time
	for i := 0; i < b.N; i++ {
		PrintProjectStructure(".", "", "", false, "text", "blue", "green", "red", []string{}, "name", "asc")
	}
}

// createLargeTestProjectStructure creates a large directory structure for benchmarking.
func createLargeTestProjectStructure(b *testing.B, root string) {
	// Create 100 directories, each containing 10 files
	for i := 0; i < 100; i++ {
		dir := filepath.Join(root, "dir"+strconv.Itoa(i))
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			b.Fatalf("Failed to create directory: %v", err)
		}

		for j := 0; j < 10; j++ {
			file := filepath.Join(dir, "file"+strconv.Itoa(j))
			f, err := os.Create(file)
			if err != nil {
				b.Fatalf("Failed to create file: %v", err)
			}
			f.Close()
		}
	}
}

package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Config holds the flag values
type Config struct {
	DirPath       string
	OutputPath    string
	ExtFilter     string
	IncludeHidden bool
}

// HandleFlags
func HandleFlags(config Config) {
	PrintProjectStructure(config.DirPath, config.OutputPath, config.ExtFilter, config.IncludeHidden)
}

// PrintProjectStructure prints the directory structure of the given root directory.
// It always prints the structure to the console and writes to the output file if provided.
func PrintProjectStructure(root string, outputFile string, extFilter string, includeHidden bool) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	rootName := filepath.Base(absRoot)
	output := fmt.Sprintf("%s/\n", rootName)
	output += getTreeOutput(absRoot, "", extFilter, includeHidden)

	fmt.Print(output)

	if outputFile != "" {
		writeToFile(output, outputFile)
	}
}

// writeToFile writes the output to the specified file
func writeToFile(output, outputFile string) {
	absOutputFile, err := filepath.Abs(outputFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}
	err = os.WriteFile(absOutputFile, []byte(output), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

// getTreeOutput returns the directory tree structure as a string.
func getTreeOutput(currentDir string, prefix string, extFilter string, includeHidden bool) string {
	var output string

	dir, err := os.Open(currentDir)
	if err != nil {
		return output
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return output
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name() < entries[j].Name()
	})

	for i, entry := range entries {
		// Skip hidden files/directories (those starting with ".")
		if !includeHidden && strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		isLast := i == len(entries)-1

		if entry.IsDir() {
			output += fmt.Sprintf("%s%s/\n", prefix+getTreePrefix(isLast), entry.Name())
			output += getTreeOutput(filepath.Join(currentDir, entry.Name()), prefix+getIndent(isLast), extFilter, includeHidden)
		} else {
			if extFilter == "" || strings.HasSuffix(entry.Name(), extFilter) {
				output += fmt.Sprintf("%s%s\n", prefix+getTreePrefix(isLast), entry.Name())
			}
		}
	}

	return output
}

// getTreePrefix returns the tree prefix for the current entry.
func getTreePrefix(isLast bool) string {
	if isLast {
		return "└── "
	}
	return "├── "
}

// getIndent returns the indentation for the current level.
func getIndent(isLast bool) string {
	if isLast {
		return "    "
	}
	return "│   "
}

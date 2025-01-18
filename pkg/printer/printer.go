package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PrintProjectStructureAndAddToDir prints the directory structure of the given root directory and writes it to the output file.
func PrintProjectStructureAndAddToDir(root string, outputFile string) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}
    
	absOutputFile, err := filepath.Abs(outputFile)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	rootName := filepath.Base(absRoot)
	output := fmt.Sprintf("%s/\n", rootName)

	output += getTreeOutput(absRoot, "")

	if absOutputFile != "" {
		err := os.WriteFile(absOutputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	} else {
		fmt.Print(output)
	}

	fmt.Println(output)
}

// PrintProjectStructure prints the folder structure starting from the specified directory.
func PrintProjectStructure(root string) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	rootName := filepath.Base(absRoot)
	fmt.Printf("%s/\n", rootName)
	fmt.Print(getTreeOutput(absRoot, ""))
}

// getTreeOutput returns the directory tree structure as a string.
func getTreeOutput(currentDir string, prefix string) string {
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
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		isLast := i == len(entries)-1

		if entry.IsDir() {
			output += fmt.Sprintf("%s%s/\n", prefix+getTreePrefix(isLast), entry.Name())
			output += getTreeOutput(filepath.Join(currentDir, entry.Name()), prefix+getIndent(isLast))
		} else {
			output += fmt.Sprintf("%s%s\n", prefix+getTreePrefix(isLast), entry.Name())
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
package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PrintProjectStructure prints the folder structure starting from the current directory.
func PrintProjectStructure() {
	root, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	rootName := filepath.Base(root)
	fmt.Printf("%s/\n", rootName)

	printTree(root, "")
}

// printTree prints the directory tree structure.
func printTree(currentDir string, prefix string) {
	dir, err := os.Open(currentDir)
	if err != nil {
		return
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1) // Read all entries
	if err != nil {
		return
	}

	// Sort entries alphabetically
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
			fmt.Printf("%s%s/\n", prefix+getTreePrefix(isLast), entry.Name())
			// Recursively print the contents of the directory
			printTree(filepath.Join(currentDir, entry.Name()), prefix+getIndent(isLast))
		} else {
			fmt.Printf("%s%s\n", prefix+getTreePrefix(isLast), entry.Name())
		}
	}
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

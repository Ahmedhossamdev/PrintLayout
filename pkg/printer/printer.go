package printer

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// PrintProjectStructure prints the folder structure starting from the specified directory.
func PrintProjectStructure(root string) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	rootName := filepath.Base(absRoot)
	fmt.Printf("%s/\n", rootName)

	printTree(absRoot, "")
}

// printTree prints the directory tree structure.
func printTree(currentDir string, prefix string) {
	dir, err := os.Open(currentDir)
	if err != nil {
		return
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return
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
			fmt.Printf("%s%s/\n", prefix+getTreePrefix(isLast), entry.Name())
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

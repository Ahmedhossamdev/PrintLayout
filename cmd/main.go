package main

import (
	"PrintLayout/pkg/printer"
	"os"
)

func main() {
	// TODO: Add advanced command-line options (depth, hidden files, etc, ignore files, etc, ignore files that in .gitignore)
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	printer.PrintProjectStructure(root)
}

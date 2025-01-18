package main

import (
	"PrintLayout/pkg/printer"
	"os"
)

func main() {

	// TODO: Add advanced command-line options (depth, hidden files, etc, ignore files, etc, ignore files that in .gitignore)

	// Get the root directory 
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	// the output file
	outputFile := ".";
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}
    
	if outputFile == "." {
		printer.PrintProjectStructure(root)
	} else {
		printer.PrintProjectStructureAndAddToDir(root, outputFile)
	}
}
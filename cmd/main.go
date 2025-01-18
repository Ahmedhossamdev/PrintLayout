package main

import (
	"PrintLayout/pkg/printer"
	"flag"
)

func main() {
	config := printer.Config{}

	flag.StringVar(&config.DirPath, "dir", ".", "Directory path to print the structure of")
	flag.StringVar(&config.OutputPath, "output", "", "Output file path")

	flag.Parse()

	printer.PrintProjectStructure(config.DirPath, config.OutputPath)
}

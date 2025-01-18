package main

import (
	"PrintLayout/pkg/printer"
	"flag"
)

func main() {
	config := printer.Config{}

	flag.StringVar(&config.DirPath, "dir", ".", "Directory path to print the structure of")
	flag.StringVar(&config.OutputPath, "output", "", "Output file path")
	flag.StringVar(&config.ExtFilter, "ext", "", "File extension filter (e.g., .go, .js)")

	flag.Parse()

	printer.PrintProjectStructure(config.DirPath, config.OutputPath, config.ExtFilter)
}

package main

import (
	"PrintLayout/pkg/printer"
	"flag"
)

func main() {
	config := printer.Config{}

	// Define flags here:
	flag.StringVar(&config.DirPath, "dir", ".", "Directory path to print the structure of")
	flag.StringVar(&config.OutputPath, "output", "", "Output file path")
	flag.StringVar(&config.ExtFilter, "ext", "", "File extension filter (e.g., .go, .js)")
	flag.BoolVar(&config.NoColor, "no-color", false, "Disable colorized output")
	flag.StringVar(&config.OutputFormat, "format", "text", "Output format (text, json, xml, yaml)")
	flag.StringVar(&config.DirColor, "dir-color", "blue", "Color for directories (e.g., blue, green, red)")
	flag.StringVar(&config.FileColor, "file-color", "green", "Color for files (e.g., yellow, cyan, magenta)")
	flag.StringVar(&config.ExecColor, "exec-color", "red", "Color for executables (e.g., red, green, blue)")
	flag.StringVar(&config.SortBy, "sort-by", "name", "Sort by 'name', 'size', or 'time'")
	flag.StringVar(&config.Order, "order", "asc", "Sort order 'asc' or 'desc'")
	flag.BoolVar(&config.IncludeHidden, "hidden", false, "Include hidden files and directories")

	// Add --exclude flag to specify exclusion patterns
	flag.Func("exclude", "Exclude files/directories matching the pattern (can be specified multiple times)", func(pattern string) error {
		config.ExcludePatterns = append(config.ExcludePatterns, pattern)
		return nil
	})

	// Parse flags
	flag.Parse()

	printer.PrintProjectStructure(
		config.DirPath,
		config.OutputPath,
		config.ExtFilter,
		!config.NoColor,
		config.OutputFormat,
		config.DirColor,
		config.FileColor,
		config.ExecColor,
		config.ExcludePatterns,
		config.SortBy,
		config.Order,
		config.IncludeHidden,
	)
}

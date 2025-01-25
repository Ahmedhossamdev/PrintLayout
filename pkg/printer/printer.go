package printer

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
)

// Config holds the flag values
type Config struct {
	DirPath         string
	OutputPath      string
	ExtFilter       string
	NoColor         bool
	OutputFormat    string
	DirColor        string
	FileColor       string
	ExecColor       string
	ExcludePatterns []string
	SortBy          string // "name", "size", "time"
	Order           string // "asc", "desc"
}

var colorMap = map[string]color.Attribute{
	"black":   color.FgBlack,
	"red":     color.FgRed,
	"green":   color.FgGreen,
	"yellow":  color.FgYellow,
	"blue":    color.FgBlue,
	"magenta": color.FgMagenta,
	"cyan":    color.FgCyan,
	"white":   color.FgWhite,
}

// getColorFunc returns a color function based on the color name
func getColorFunc(colorName string) func(a ...interface{}) string {
	if attr, ok := colorMap[colorName]; ok {
		return color.New(attr).SprintFunc()
	}
	return fmt.Sprint // Default to no color if the color name is invalid
}

// HandleFlags processes the configuration and prints the directory structure.
func HandleFlags(config Config) {
	PrintProjectStructure(
		config.DirPath,
		config.OutputPath,
		config.ExtFilter,
		!config.NoColor, config.OutputFormat,
		config.DirColor,
		config.FileColor,
		config.ExecColor,
		config.ExcludePatterns,
		config.SortBy,
		config.Order)
}

// PrintProjectStructure prints the directory structure of the given root directory.
func PrintProjectStructure(
	root string,
	outputFile string,
	extFilter string,
	useColor bool,
	format string,
	dirColorName string,
	fileColorName string,
	execColorName string,
	excludePatterns []string,
	sortBy string,
	order string) {
	absRoot, err := filepath.Abs(root)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	if format == "text" {
		dirCount, fileCount := getTreeOutput(absRoot, extFilter, useColor, dirColorName, fileColorName, execColorName, excludePatterns, sortBy, order)
		fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
	} else {
		tree := buildTree(absRoot, extFilter, excludePatterns, sortBy, order)
		var output string
		switch format {
		case "json":
			data, _ := json.MarshalIndent(tree, "", "  ")
			output = string(data)
		case "xml":
			data, _ := xml.MarshalIndent(tree, "", "  ")
			output = string(data)
		case "yaml":
			data, _ := yaml.Marshal(tree)
			output = string(data)
		default:
			fmt.Println("Unsupported format:", format)
			return
		}

		fmt.Print(output)

		if outputFile != "" {
			writeToFile(output, outputFile)
		}
	}
}

func getTreeOutput(root string, extFilter string, useColor bool, dirColorName string, fileColorName string, execColorName string, excludePatterns []string, sortBy string, order string) (int, int) {
	dirCount := 0
	fileCount := 0

	dirColorFunc := getColorFunc(dirColorName)
	fileColorFunc := getColorFunc(fileColorName)
	execColorFunc := getColorFunc(execColorName)

	var traverse func(string, string) error
	traverse = func(currentDir string, prefix string) error {
		dir, err := os.Open(currentDir)
		if err != nil {
			return err
		}
		defer dir.Close()

		entries, err := dir.Readdir(-1)
		if err != nil {
			return err
		}

		// Sort entries based on the specified criteria and order
		sortEntries(entries, sortBy, order)

		for i, entry := range entries {
			if strings.HasPrefix(entry.Name(), ".") {
				continue
			}

			if isExcluded(entry.Name(), excludePatterns) {
				if entry.IsDir() {
					continue
				}
				continue
			}

			isLast := i == len(entries)-1

			if entry.IsDir() {
				dirCount++
				if useColor {
					fmt.Printf("%s%s/\n", prefix+getTreePrefix(isLast), dirColorFunc(entry.Name()))
				} else {
					fmt.Printf("%s%s/\n", prefix+getTreePrefix(isLast), entry.Name())
				}

				err := traverse(filepath.Join(currentDir, entry.Name()), prefix+getIndent(isLast))
				if err != nil {
					return err
				}
			} else {
				if extFilter == "" || strings.HasSuffix(entry.Name(), extFilter) {
					fileCount++
					if useColor {
						info, err := os.Stat(filepath.Join(currentDir, entry.Name()))
						if err != nil {
							fmt.Printf("%s%s\n", prefix+getTreePrefix(isLast), entry.Name())
						} else if isExecutable(info) {
							fmt.Printf("%s%s\n", prefix+getTreePrefix(isLast), execColorFunc(entry.Name()))
						} else {
							fmt.Printf("%s%s\n", prefix+getTreePrefix(isLast), fileColorFunc(entry.Name()))
						}
					} else {
						fmt.Printf("%s%s\n", prefix+getTreePrefix(isLast), entry.Name())
					}
				}
			}
		}

		return nil
	}

	fmt.Printf("%s/\n", filepath.Base(root))
	err := traverse(root, "")
	if err != nil {
		fmt.Println("Error traversing directory:", err)
	}

	return dirCount, fileCount
}

// sortEntries sorts the entries based on the specified criteria and order
func sortEntries(entries []os.FileInfo, sortBy string, order string) {
	switch sortBy {
	case "name":
		sort.Slice(entries, func(i, j int) bool {
			if order == "asc" {
				return entries[i].Name() < entries[j].Name()
			}
			return entries[i].Name() > entries[j].Name()
		})
	case "size":
		sort.Slice(entries, func(i, j int) bool {
			if order == "asc" {
				return entries[i].Size() < entries[j].Size()
			}
			return entries[i].Size() > entries[j].Size()
		})
	case "time":
		sort.Slice(entries, func(i, j int) bool {
			if order == "asc" {
				return entries[i].ModTime().Before(entries[j].ModTime())
			}
			return entries[i].ModTime().After(entries[j].ModTime())
		})
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

// Node represents a directory or file in the tree structure
type Node struct {
	Name     string  `json:"name" xml:"name"`
	IsDir    bool    `json:"is_dir" xml:"is_dir"`
	Children []*Node `json:"children,omitempty" xml:"children,omitempty"`
}

// buildTree constructs a tree of Nodes from the directory structure
func buildTree(currentDir string, extFilter string, excludePatterns []string, sortBy string, order string) *Node {
	dir, err := os.Open(currentDir)
	if err != nil {
		return nil
	}
	defer dir.Close()

	entries, err := dir.Readdir(-1)
	if err != nil {
		return nil
	}

	// Sort entries based on the specified criteria and order
	sortEntries(entries, sortBy, order)

	node := &Node{
		Name:  filepath.Base(currentDir),
		IsDir: true,
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		// Check if the entry matches any exclusion pattern
		if isExcluded(entry.Name(), excludePatterns) {
			continue
		}

		if entry.IsDir() {
			child := buildTree(filepath.Join(currentDir, entry.Name()), extFilter, excludePatterns, sortBy, order)
			if child != nil {
				node.Children = append(node.Children, child)
			}
		} else if extFilter == "" || strings.HasSuffix(entry.Name(), extFilter) {
			node.Children = append(node.Children, &Node{
				Name:  entry.Name(),
				IsDir: false,
			})
		}
	}

	return node
}

// isExecutable checks if a file is executable
func isExecutable(entry os.FileInfo) bool {
	return entry.Mode()&0111 != 0 // Check executable bits
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

// isExcluded checks if a file/directory matches any of the exclusion patterns
func isExcluded(name string, excludePatterns []string) bool {
	for _, pattern := range excludePatterns {
		matched, err := filepath.Match(pattern, name)
		if err != nil {
			fmt.Printf("Invalid exclude pattern: %s\n", pattern)
			continue
		}
		if matched {
			return true
		}
	}
	return false
}

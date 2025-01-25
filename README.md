# PrintLayout

![PrintLayout Logo](assets/printlayoutlogo.webp)

PrintLayout is a command-line tool that prints the directory structure of a specified folder in a tree format. It is designed to be simple, fast, and customizable.

## Installation

### Option 1: Go Install

To install PrintLayout, run:

```bash
go install github.com/Ahmedhossamdev/PrintLayout/cmd/main.go
```

This will install the printlayout executable to your $GOPATH/bin directory. Make sure $GOPATH/bin is in your PATH environment variable.

### Option 2: Download Pre-Built Binaries

Download the pre-built binary for your operating system from the [Releases page](#).

#### Linux/macOS

1. Download the binary for your platform (e.g., `printlayout-linux-amd64` or `printlayout-darwin-amd64`).
2. Make the binary executable:
   ```bash
   chmod +x printlayout-linux-amd64
   ```
3. Move the binary to a directory in your `PATH` (e.g., `/usr/local/bin`):
   ```bash
   sudo mv pr-linux-amd64 /usr/local/bin/printlayout
   ```
4. Run the program:
   ```bash
      pr -dir /path/to/your/folder or pr to print the current dir
   ```

#### Windows

1. Download the binary for your platform (e.g., `printlayout-windows-amd64.exe`).
2. Move the binary to a directory in your `PATH` (e.g., `C:\Windows\System32`).

## Usage

### Print the Directory Structure

To print the directory structure of the current folder:

```bash
pr
```

To print the directory structure of a specific folder:****

```bash
pr -dir /path/to/your/folder
```

To print the directory structure and filter by a specific file extension (e.g., .go, .js, .rb):

```bash
pr -ext .js
```

To print the directory structure of a specific folder, filter by a specific file extension, and save the output to a file:

```bash
pr -dir /path/to/your/folder -ext .rb -output output.txt
```

### Example Output

```bash
$ pr -dir /path/to/your/folder
printLayout/
├── cmd/
│   └── main.go
├── go.mod
├── internal/
│   └── utils/
│       └── utils.go
└── pkg/
    └── printer/
        ├── printer.go
        └── printer_test.go
```

# `pr` Command-Line Flags

## Basic Flags

### `--dir`
- **Description**: Specifies the directory to print the structure of.
- **Default**: Current directory (`.`)
- **Example**: `pr --dir /path/to/your/folder`

### `--ext`
- **Description**: Filters files by a specific extension (e.g., `.go`, `.js`)
- **Default**: No filter (all files are included)
- **Example**: `pr --ext .go`

### `--output`
- **Description**: Saves the output to a file instead of printing it to the terminal
- **Default**: Output is printed to the terminal
- **Example**: `pr --output output.txt`

### `--no-color`
- **Description**: Disables colored output
- **Default**: Colors are enabled
- **Example**: `pr --no-color`

## Sorting Flags

### `--sort-by`
- **Description**: Specifies the sorting criteria
- **Options**:
  - `name`: Sort by file/directory name
  - `size`: Sort by file size
  - `time`: Sort by modification time
- **Default**: `name`
- **Example**: `pr --sort-by size`

### `--order`
- **Description**: Specifies the sorting order
- **Options**:
  - `asc`: Ascending order
  - `desc`: Descending order
- **Default**: `asc`
- **Example**: `pr --sort-by time --order desc`

## Exclusion Flags

### `--exclude`
- **Description**: Excludes files or directories matching the specified pattern (can be used multiple times)
- **Default**: No exclusions
- **Example**: `pr --exclude "*.log" --exclude "tmp/*"`

## Output Format Flags

### `--format`
- **Description**: Specifies the output format
- **Options**:
  - `text`: Plain text format (default)
  - `json`: JSON format
  - `xml`: XML format
  - `yaml`: YAML format
- **Default**: `text`
- **Example**: `pr --format json`

## Color Customization Flags

### `--dir-color`
- **Description**: Specifies the color for directories
- **Options**: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`
- **Default**: `blue`
- **Example**: `pr --dir-color green`

### `--file-color`
- **Description**: Specifies the color for files
- **Options**: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`
- **Default**: `green`
- **Example**: `pr --file-color yellow`

### `--exec-color`
- **Description**: Specifies the color for executable files
- **Options**: `black`, `red`, `green`, `yellow`, `blue`, `magenta`, `cyan`, `white`
- **Default**: `red`
- **Example**: `pr --exec-color magenta`

## Help Flag

### `--help` or `-h`
- **Description**: Displays the help message with all available flags and usage examples
- **Example**: `pr --help`

## Combined Flags Example

```bash
pr --dir /path/to/your/folder --ext .go --sort-by size --order desc --output output.txt --no-color
```

## Development

### Run the Project in Development

To run the project during development without installing it:

```bash
go run ./cmd/main.go
```

### Run Tests

To run the tests:

```bash
go test -v ./...
```

### Build the Project

To build the project:

```bash
go build -o printlayout ./cmd/main.go
```

This will create an executable named printlayout in your project directory.

## Future Improvements (TODOs)

Here are some ideas for future improvements to the project:

1. Advanced Command-Line Options:

   - [ ] Add support for limiting the depth of the directory tree (e.g., --depth 2).
   - [ ] an option to include hidden files and directories (e.g., --hidden).
   - [ ] an option to ignore files and directories listed in .gitignore (e.g., --ignore-gitignore).

2. Customizable Output:

   - [ ] support for customizing the tree symbols (e.g., --symbols=ascii for ASCII-only output).
   - [x] support for exporting the directory structure to a file (e.g., --output tree.txt).

3. Performance Improvements:
   - [ ] Optimize the directory traversal for large directories.
   - [ ] Add support for parallel processing of directories.

## Contributing

Contributions are welcome! If you'd like to contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Make your changes and write tests if applicable.
4. Submit a pull request with a detailed description of your changes.


## Acknowledgments

- Built with Go for simplicity and performance.
- Inspired by GNU Tree: Incorporating all the features of GNU Tree while innovating with new functionalities to enhance the user experience.

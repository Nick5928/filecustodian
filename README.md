# "File Custodian" : A terminal based directory analyzer application
## Description
File Custodian analyzes disk usage across your directories, showing how much space each file type and programming language consumes. It helps you quickly spot large or unnecessary files and understand how your projects and folders are organized.

## Motivation
As a gamer with a lot of installed games (and friends in the same boat), I kept running into a simple problem: my storage would fill up, and I had no idea which game folders were actually using the most space. File Explorer can show total size, but it’s painful to see the sizes of many child directories at once. I built this tool to scan a parent directory, find the largest child directories, and display useful details about each one, so it’s easy to track down what’s eating disk space.
## Quick Start
### Install filecustodian using the Go toolchain
```bash
# Install File Custodian
go install github.com/Nick5928/filecustodian

# Build cli
go build -o filecustodian
```
## Usage
- `calcsize` - Calculates the size of directory and provides usage for top 10 biggest child directories
- `help` - Lists commands and respective options
## Contributing
### Clone the repo
```bash
git clone https://github.com/nick5928/filecustodian
cd filecustodian
```

### Build the binary
```bash
go build
```

### Submit a pull request
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
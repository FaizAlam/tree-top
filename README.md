# tree-top

**tree-top** is a fast, interactive, terminal-based directory explorer built in Go. It provides a visual, `htop`-like interface for navigating nested filesystems, inspecting metadata, and even previewing file contents—all without leaving your terminal.

![tree-top-gif](https://github.com/user-attachments/assets/3079307f-c3f0-4bc5-8be0-842dbf644a2e)

## Features

- **Nested Tree View**: Expand and collapse folders to explore directories in a structured tree.
- **Keyboard Navigation**: Use `↑`/`↓` to move, `→` to expand, `←` to collapse, and `q` to quit.
- **Persistent Context**: Once expanded, directories stay open as you navigate deeper.
- **Metadata Panel**: View file/folder details—permissions, type, size, timestamps, and full path—in a side panel.
- **Directory Highlighting**: Folders are styled with a colored text for quick recognition.
- **Custom Root**: Optionally specify a directory to explore: `tree-top -dir=[path]`.

## Installation

### Prerequisites

- Go 1.20 or later installed
- A terminal emulator that supports ANSI colors

### Install via `go install`

```bash
go install github.com/faizalam/tree-top/cmd/tree-top@latest
```

This places the `tree-top` binary in your `$GOBIN`. Run it with:
```bash
tree-top
```

### Build from Source

```bash
git clone https://github.com/faizalam/tree-top.git
cd tree-top
go mod tidy
go build -o tree-top ./cmd/tree-top
./tree-top
```

## Usage

Launch the tool in any directory:

```bash
tree-top
```

or specify a custom root directory:
```bash
tree-top -dir=/path/to/dir
```

- **Navigate**: `↑`/`↓` to move selection
- **Expand folder**: `→`
- **Collapse folder**: `←`
- **Quit**: `q`

## Development Setup

1. **Clone the repo**
   ```bash
    git clone https://github.com/faizalam/tree-top.git
    cd tree-top
    ```  
2. **Install dependencies**
   ```bash
    go mod tidy
    ```  
3. **Run the app**
   ```bash
    go run ./cmd/tree-top
    ```  
4. **Lint & Format**
   ```bash
    golangci-lint run
    go fmt ./...
    ```  

## Configuration

Currently there are no configurable options. Future releases may include:

- Custom keybindings via `config.yaml`
- Theme support (light/dark modes)

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/YourFeature`
3. Commit your changes
4. Open a pull request

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.

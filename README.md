# calculator

[![CI](../../actions/workflows/ci.yaml/badge.svg)](../../actions/workflows/ci.yaml)

A Command-Line Interface (CLI) Calculator Written in Go.

## Installation

### Using Nix

If you have Nix with flakes enabled:

```bash
nix run git+https://github.com/stefankuehnel/calculator -- <args>
```

Or install it to your profile:

```bash
nix profile add git+https://github.com/stefankuehnel/calculator
```

### Using Docker

```bash
docker run --rm ghcr.io/stefankuehnel/calculator <args>
```

### Using Go

```bash
go install github.com/stefankuehnel/calculator@latest
```

### Building from Source

```bash
git clone https://github.com/stefankuehnel/calculator.git
cd calculator
go build
```

## Development

This project uses [Task](https://taskfile.dev) as a task runner.

### Available Tasks

```bash
# Run default tasks (lint, build and test)
task

# Initialize project
task init

# Run project
task run -- <args>

# Build project
task build

# Deploy project
task deploy

# Format project
task format

# Lint project
task lint

# Test project
task test

# Test project with coverage
task test:coverage

# Clean project
task clean
```

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
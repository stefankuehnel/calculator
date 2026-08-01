# calculator

[![CI](../../actions/workflows/ci.yaml/badge.svg)](../../actions/workflows/ci.yaml)
[![CD](../../actions/workflows/cd.yaml/badge.svg)](../../actions/workflows/cd.yaml)

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
# Run default task
task

# Initialize project
task init

# Generate code
task codegen

# Format project
task format

# Lint project
task lint

# Test project
task test

# Test project with coverage
task test:coverage

# Build project
task build

# Load project
task load

# Run project
task run -- <args>

# Run project locally with Workflow
task run:workflow:local

# Validate project
task validate

# Deploy project
task deploy

# Clean project
task clean

# List all available tasks
task --list-all --sort=none
```

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.
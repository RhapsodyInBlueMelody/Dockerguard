# Dockguard

**Dockguard** is a lightweight command-line tool written in Go that checks Dockerfiles for potential configuration issues and security concerns.

## Overview

Dockerfiles are an important part of building container images, but small configuration mistakes can introduce unnecessary risks. Dockguard aims to help developers identify selected issues in their Dockerfiles through simple, readable warnings.

This project is being developed as a learning exercise in Go programming, command-line tooling, and container security best practices.

## Features

- Detects `FROM` instructions that explicitly use the `latest` image tag.
- Reports warnings with Dockerfile line numbers.
- Provides a foundation for additional Dockerfile checks.

> **Note:** Dockguard is under active development. Features and detection coverage may change.

## Requirements

- Go
- A Dockerfile to scan

## Usage

Run Dockguard from the project directory:

```bash
go run . Dockerfile
```

Replace `Dockerfile` with the path to the file you want to scan.

## Testing

Run the test suite with:

```bash
go test ./...
```

## Limitations

Dockguard uses simple checks and does not fully parse or validate Dockerfiles. It may miss certain issues or produce inaccurate results.

Dockguard is not a substitute for a comprehensive container security review.

## Development

Contributions, bug reports, and suggestions are welcome. Please ensure that tests pass before submitting changes.

## License

No license has been specified yet.

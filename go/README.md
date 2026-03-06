# My Routine For Starting a Go Project

All commands should be run from the `go/` directory. Adjust paths as needed if you change the structure.

## Initialize a project

Initialize the module (replace with your real repository path):

```bash
go mod init github.com/<github-username>/<project-name>/go
```

## Manage dependencies

```bash
# Add a dependency
go get <package-path>

# Clean and sync go.mod/go.sum
go mod tidy
```

## Linting, type checking, and formatting

Go type checking is built into `go build` / `go test`. For linting, this template uses [golangci-lint](https://github.com/golangci/golangci-lint).

Install `golangci-lint` (pin a version for consistent CI results):

```bash
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v2.10.1
```

Run checks:

```bash
golangci-lint run
```

Lint config lives in `.golangci.yml`.

## Project structure

This template follows the common Go layout:

- `cmd/`: application entrypoints (`main` packages)
- `internal/`: private code, importable only within this module
- `pkg/`: reusable library code for external usage (optional)

## Testing

This template uses [Testify](https://github.com/stretchr/testify) for assertions and mocking.

**Naming convention**: Go test files should be named `*_test.go` and test functions should be named `TestXxx`.

To run tests:

```bash
go test ./...
go test -v ./...
go test -cover ./...
```

## Dockerfile (Podman/Docker)

To build and run the Docker image:

```bash
podman build -t localhost/my-go-app .
podman run --rm localhost/my-go-app
```

Dockerfile template lives in `Dockerfile` and `.dockerignore`.

Notes:
- This starter app is CLI-style, so `EXPOSE` is intentionally not enabled by default.
- Add `EXPOSE <port>` only when your app actually listens on that port.
- The container runs as non-root `UID:GID 65532:65532`.

## Live reloading (optional)

Use [Air](https://github.com/air-verse/air) for local live reload:

```bash
go install github.com/air-verse/air@latest
air init
air
```

## VS Code integration

Install the **Go** extension (`golang.go`) for linting, formatting, and test integration.

VS Code settings lives in `../.vscode/settings.json`.

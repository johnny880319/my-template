# AGENTS.md

## Purpose

This repository is a starter kit for new codebases. Each language folder should stay minimal, runnable, and easy to copy into a new project.

## Repository layout

- `README.md`: shared setup flow (git, basic project bootstrap)
- `go-codebase/`: Go starter template
- `python-codebase/`: Python starter template
- `rust-codebase/`: Rust starter template
- `typescript-codebase/`: TypeScript starter template

## Global rules for agents

- Keep templates small and practical; avoid adding product-specific logic.
- Keep docs command-first and concise.
- When a command or file tree changes, update the corresponding README in the same folder.
- Prefer stable tooling and reproducible setup instructions.
- Prefer examples that run from each language directory (`cd go-codebase`, `cd python-codebase`, `cd rust-codebase`, `cd typescript-codebase`) for consistency.
- Keep runtime demos explicit (for example, print a success message) so container smoke tests are easy.

## Template expansion rules

- Prefer adding small, composable templates over one large "do everything" template.
- Add templates only when they are reusable across multiple new projects.
- For topics with many valid approaches, provide a minimal baseline plus optional variants.
- Keep "opinionated defaults" explicit in docs so users can swap tools quickly.

## Monorepo wiring conventions

- Keep CI workflows language-scoped with `paths` filters and per-language working directories.
- Keep hook commands language-scoped in `lefthook.yml` using each language folder as `root`.
- Prefer short local commands in language READMEs (assume the user runs commands after `cd` into that folder).

## Go template conventions

- Keep entrypoints in `go-codebase/cmd/<app-name>/`.
- Keep private/internal code in `go-codebase/internal/`.
- Put optional reusable libraries in `go-codebase/pkg/`.
- Keep package names lowercase and simple.
- Keep tests next to implementation files with `*_test.go` naming.
- Validate changes from `go-codebase/` with:
  - `go test ./...`
  - `go test -cover ./...`
  - `go fmt ./...`
  - `golangci-lint run`

## Python template conventions

- Use `uv` for dependency and environment management.
- Keep source under `python-codebase/src/` and tests under `python-codebase/tests/`.
- Validate changes from `python-codebase/` with:
  - `uv sync`
  - `uv run ruff check .`
  - `uv run ruff format .`
  - `uv run pyright .`
  - `uv run pytest`

## TypeScript template conventions

- Use `pnpm` for package management and scripts.
- Keep source under `typescript-codebase/src/` and tests under `typescript-codebase/tests/`.
- Use Biome as the default formatter+linter baseline for starter simplicity.
- Validate changes from `typescript-codebase/` with:
  - `pnpm install`
  - `pnpm run lint`
  - `pnpm run typecheck`
  - `pnpm run test`

## Documentation checklist

Before finishing a change:

- Ensure commands in README are executable as written.
- Ensure folder trees in README match the real files.
- Remove stale examples and dead references.
- Keep wording concise and avoid repeating the same instruction in multiple sections.

## Container template policy

- Use `Dockerfile` as the default filename (compatible with both Podman and Docker).
- Keep container templates language-specific (`go-codebase/`, `python-codebase/`) instead of a shared generic file.
- Prefer one practical baseline Dockerfile per language; add variants only when clearly needed.
- Keep runtime assumptions explicit (entrypoint, env vars, exposed ports) and easy to override.

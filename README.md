# my-template

This repository is a starter kit for new codebases with Go, Python, Rust, and TypeScript templates.

## Getting started

```bash
git clone git@github.com:johnny880319/my-template.git
cd my-template
```

## Post-copy setup

Install Lefthook if you do not have it:

```bash
go install github.com/evilmartians/lefthook@latest
```

Install Git hooks:

```bash
lefthook install
```

Hook config lives in `lefthook.yml`.

## GitHub Actions CI

This template includes language-specific CI workflows:

- `python-ci.yml`: runs `ruff`, `pyright`, and `pytest` for `python-codebase/`
- `go-ci.yml`: runs `golangci-lint` and `go test` for `go-codebase/`
- `typescript-ci.yml`: runs `biome`, `tsc`, and `vitest` for `typescript-codebase/`

## Conventional workflow

1. [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
2. [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
3. [Semantic Versioning](https://semver.org/)
4. [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)

## VS Code integration

Install **GitLens** and **Git Graph** extensions for advanced Git UI.
Install the **Jupyter** extension for notebook support.

VS Code settings lives in `.vscode/settings.json`.

# My routine of starting a new project

## Initialize git repository

Command to initialize a git repository:

```bash
git init
```

Set up git user name and email:

```bash
git config --local user.name "Your Name"
git config --local user.email "your.email@example.com"
```

## Create essential files

```bash
touch README.md Makefile .gitignore 
```

VS Code users:

```bash
mkdir -p .vscode && touch .vscode/settings.json
```

## Conventional workflow

Follow these industry standards for collaboration and versioning:

1. [GitHub Flow](https://docs.github.com/en/get-started/quickstart/github-flow)
2. [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
3. [Semantic Versioning](https://semver.org/)
4. [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)

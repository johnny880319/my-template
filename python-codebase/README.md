# My routine of starting a python project

All commands should be run from the `python-codebase/` directory. Adjust paths as needed if you change the structure.

## Initialize a project

[UV](https://astral.sh/uv) is a blazingly fast Python package manager written in Rust, replacing pip and virtualenv.

Install on Linux/macOS if you do not have it already (check docs for Windows):

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

Initialize a project with current settings:

```bash
uv sync
```

## Manage dependencies

Manage dependencies:

```bash
uv add <package-name>
uv remove <package-name>
```

Manage development dependencies (excluded from production):

```bash
uv add --dev <package-name>
uv remove --dev <package-name>
```

Use `uv sync` to install dependencies in `pyproject.toml` and update `uv.lock` (for local dev). Add `--locked` to strictly install from the lockfile without updating it (for CI/CD).

```bash
uv sync
uv sync --locked
```

## Linting, type checking and formatting

We use [ruff](https://github.com/charliermarsh/ruff) and [pyright](https://github.com/microsoft/pyright) as linters and type checkers (pyright[nodejs] includes a standalone Node environment for stability):

To run these tools from the command line:

```bash
uv run ruff check --fix <file_or_dir>
uv run ruff format <file_or_dir>
uv run pyright <file_or_dir>
```

Lint config lives in `pyproject.toml`.

## Project structure

This project uses the standard Python `src` layout.

* **`src/`**: Application code. Prevents import errors and ensures test reliability.
* **`tests/`**: Contains all unit and integration tests.
* **`./`**: Configuration files and documentation.

## Testing

This template uses [`pytest`](https://docs.pytest.org/) for unit testing.

**Naming conventions:** `pytest` automatically discovers tests. Ensure your test files are named `test_*.py` (or `*_test.py`) and your test functions start with `test_`.

To run the tests:

```bash
uv run pytest
```

## Dockerfile (Podman/Docker)

To build and run the Docker image:

```bash
podman build -t localhost/my-python-app .
podman run --rm localhost/my-python-app
```

Dockerfile template lives in `Dockerfile` and `.dockerignore`.

Note:
- The Dockerfile uses `ENTRYPOINT ["python", "-m"]` + `CMD ["my_package.main"]`.
- Override `CMD` to run a different module without rebuilding the image.
- Dependencies are isolated under `/opt/.venv`, separate from app code in `/app`.
- The container runs as non-root `UID:GID 65532:65532`.

## Jupyter Notebook support

Install [ipykernel](https://ipykernel.readthedocs.io/) for Jupyter Notebook support:

```bash
uv add --dev ipykernel
```

## VS Code integration

Install the **Ruff** and **Pyright** extensions for linting support.
Install the **Jupyter** extension for notebook support.

VS Code settings lives in `../.vscode/settings.json`.

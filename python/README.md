# My routine of starting a python project

## Manage environments

### Install UV (If you don't have it already)

[UV](https://astral.sh/uv) is a blazingly fast Python package manager written in Rust, replacing pip and virtualenv.

Install on Linux/macOS (check docs for Windows):

```bash
curl -LsSf https://astral.sh/uv/install.sh | sh
```

### Create a new project

Initialize a project:

```bash
uv init
```

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

Install [ruff](https://github.com/charliermarsh/ruff) and [pyright](https://github.com/microsoft/pyright) as linters and type checkers (pyright[nodejs] includes a standalone Node environment for stability):

```bash
uv add --dev ruff "pyright[nodejs]"
```

**Add** to `pyproject.toml`:

```toml
[tool.ruff]
line-length = 120
fix = true
exclude = [
    ".venv",
    "libs",
]

[tool.ruff.lint]
# select all rules except the ones that conflict with others rules or formatters.
select = ["ALL"]
ignore = [
    "D203", # incorrect-blank-line-before-class, conflicts with D211
    "D213", # multi-line-summary-second-line, conflicts with D212
    "COM812", # missing-trailing-comma, conflicts with formatters
]

[tool.pyright]
typeCheckingMode = "strict"
```

To run these tools from the command line:

```bash
uv run ruff check <file_or_dir>
uv run ruff check --fix <file_or_dir>
uv run ruff format <file_or_dir>
uv run pyright <file_or_dir>
```

### VS Code integration
Install the **Ruff** and **Pyright** extensions for editor support.

**Add** to `.vscode/settings.json`

```json
{
  "[python]": {
    "editor.defaultFormatter": "charliermarsh.ruff",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.fixAll": "explicit"
    }
  }
}
```

## Project structure

This project uses the standard Python `src` layout.

* **`src/`**: Application code. Prevents import errors and ensures test reliability.
* **`tests/`**: Contains all unit and integration tests.
* **`./`**: Configuration files and documentation.

```
my_python_project/
├── .venv/                  
├── src/
│   └── my_package/         
│       ├── __init__.py    
│       └── main.py         
├── tests/
│   ├── __init__.py         
│   └── test_main.py  
├── .gitignore                   
├── .python-version         
├── pyproject.toml
├── README.md
└── uv.lock
```

## Testing

Install [`pytest`](https://docs.pytest.org/) for unit testing.

```bash
uv add --dev pytest
```

**Naming conventions:** `pytest` automatically discovers tests. Ensure your test files are named `test_*.py` (or `*_test.py`) and your test functions start with `test_`. To run the tests, simply execute `uv run pytest`.


## Jupyter Notebook support

Install [ipykernel](https://ipykernel.readthedocs.io/) for Jupyter Notebook support:

```bash
uv add --dev ipykernel
```

### VS Code integration
Install the **Jupyter** extension for notebook support.

**Add** to `.vscode/settings.json` for output scrolling:

```json
{
  "notebook.output.scrolling": true
}
```
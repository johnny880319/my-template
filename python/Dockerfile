# syntax=docker/dockerfile:1

# Build stage
FROM docker.io/library/python:3.12-slim AS builder

COPY --from=ghcr.io/astral-sh/uv:0.10.7 /uv /bin/uv
ENV UV_COMPILE_BYTECODE=1
ENV UV_LINK_MODE=copy

ENV UV_PROJECT_ENVIRONMENT="/opt/.venv"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

WORKDIR /app

COPY pyproject.toml uv.lock README.md ./
RUN uv sync --locked --no-install-project --no-dev

COPY src ./src
RUN uv sync --locked --no-dev

# Runtime stage
FROM docker.io/library/python:3.12-slim

RUN useradd -m -u 1000 appuser
USER appuser

WORKDIR /app

COPY --from=builder --chown=appuser:appuser /opt/.venv /opt/.venv
COPY --from=builder --chown=appuser:appuser /app /app

ENV PATH="/opt/.venv/bin:$PATH"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# Replace module path if your app entrypoint changes.
ENTRYPOINT ["python", "-m"]
CMD ["my_package.main"]

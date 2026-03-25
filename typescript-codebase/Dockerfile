# syntax=docker/dockerfile:1

# Build stage
FROM docker.io/library/node:24-alpine AS builder

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

COPY src ./src
COPY tsconfig.json tsconfig.build.json ./

RUN pnpm tsc -p tsconfig.build.json

# Runtime stage
FROM docker.io/library/node:24-alpine

WORKDIR /app

RUN deluser --remove-home node && \
    addgroup -g 65532 nonroot && \
    adduser -u 65532 -G nonroot -S nonroot


COPY --from=builder --chown=65532:65532 /app/package.json ./package.json
COPY --from=builder --chown=65532:65532 /app/node_modules ./node_modules
COPY --from=builder --chown=65532:65532 /app/dist ./dist

# Run as non-root for safer defaults.
USER 65532:65532

# Replace module path if your app entrypoint changes.
ENTRYPOINT ["node"]
CMD ["dist/cli.js"]

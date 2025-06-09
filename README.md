# MediaDB

> **Under Construction**

## Development

```bash
# Running environment
docker compose -f dev-compose.yaml up -d

# Stopping environment
docker compose -f dev-compose.yaml down

# Run tests
go test ./...

# Run/rerun all tests
go test -count=1 ./...

# Build container
docker build -t mediadb .
```

## Build

```bash
# Running environment
docker compose -d --build

# Stopping environment
docker compose down
```


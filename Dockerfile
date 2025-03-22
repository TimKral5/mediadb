# Stage 1: Install Packages
FROM oven/bun:latest AS installer

WORKDIR /src
COPY ./src .

RUN bun install

# Stage 2: Build Application
FROM installer AS builder

WORKDIR /src

RUN \
  bun build \
    --compile \
    --outfile mdb \
    --target=bun-linux-x64 \
    ./index.ts

# Stage 3: Create Image
FROM ubuntu:latest

WORKDIR /app
COPY --from=builder /src/mdb .

RUN chmod +x /app/mdb

EXPOSE 3000
CMD ["/app/mdb"]

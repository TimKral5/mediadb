# Stage 1: Build application
FROM golang AS builder

WORKDIR /src
COPY . .
RUN go build .

# Stage 2: Setup image and install dependencies
FROM alpine:latest AS base
RUN apk add libc6-compat

# Stage 3: Install application binaries
FROM base
WORKDIR /app
COPY --from=builder /src/mediadb /app
RUN chmod +x /app/mediadb

EXPOSE 3000
CMD ["/app/mediadb"]

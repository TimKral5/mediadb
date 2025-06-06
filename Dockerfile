# Stage 1: Setup build dependencies
FROM golang AS build_dep
WORKDIR /src

COPY . .
RUN go mod download -x

# Stage 2: Build application
FROM build_dep AS build
RUN go build .

# Stage 3: Setup image and install dependencies
FROM alpine:latest AS base
RUN apk add libc6-compat

# Stage 4: Install application binaries
FROM base
WORKDIR /app
COPY --from=build /src/mediadb /app
RUN chmod +x /app/mediadb

EXPOSE 3000
CMD ["/app/mediadb"]

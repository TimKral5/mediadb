ARG BUILD_CONFIG=release

# Stage 1: Setup build dependencies
FROM golang AS build_dep
WORKDIR /src

COPY . .
RUN go mod download -x

COPY ./scripts/run-dev.sh /run.sh
COPY ./scripts/run-prod.sh .
RUN chmod +x /run.sh

# Stage 2: Build documentation
FROM openapitools/openapi-generator-cli AS build_docs
WORKDIR /src

COPY . .
RUN /usr/local/bin/docker-entrypoint.sh generate -i /src/openapi.yaml -g html -o /src/api-docs/html

# Stage 3: Build application
FROM build_dep AS build
COPY --from=build_docs /src/api-docs ./api-docs
RUN go build .
RUN chmod +x /src/mediadb

# Stage 4: Setup image and install dependencies
FROM alpine:latest AS base
RUN apk add libc6-compat

# Stage 5: Install application binaries
FROM base AS alpine_build
WORKDIR /app
COPY --from=build /src/mediadb /app
COPY --from=build /src/run-prod.sh /run.sh
RUN chmod +x /app/mediadb
RUN chmod +x /run.sh

# Stage 6: Select Output
FROM build AS debug_build
FROM alpine_build AS release_build

FROM ${BUILD_CONFIG}_build AS output
EXPOSE 3000
CMD ["/run.sh"]


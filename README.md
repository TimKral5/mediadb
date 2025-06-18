# MediaDB

> **Under Construction**

MediaDB is a RestAPI that provides stored metadata for media such as
movies, shows, books and audio.

## Technologies

This project has been developed using following technologies:

- [Go](https://go.dev/) and standard libraries
- [MongoDB](https://www.mongodb.com/)
- [Bitnami OpenLDAP](https://hub.docker.com/r/bitnami/openldap)
- [Docker](https://www.docker.com/)
- [RCF2307bis LDAP Schema](https://github.com/jtyr/rfc2307bis)

## Launch the Project

Running MediaDB locally can be achieved using Docker. For that simply
run `docker compose up --build -d`. This will build and launch the
project in a container.

Alternatively, the project can be run outside of a container, which
requires `docker compose up -d` to be run to launch all containers.
Then, run `source env.sh` in order to setup the required environment
variables. Now, to run the application itself, simply run `go run .`.

## Testing

The unit tests can be run using `go test ./...`.

If the `BUILD_CONFIG` argument in the `compose.yaml` is set to
`debug`, the tests are being run on each startup of the container.


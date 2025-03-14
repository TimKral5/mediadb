# MediaDB

MediaDB is a RestAPI that serves metadata for media such as movies,
shows and books.

## Technologies

The API itself is a basic TypeScript application that uses
**ExpressJS** to serve the data.

The data is hosted on a **MongoDB** instance, while logs and metrics
are sent to a **Loki** and a **Prometheus** instance.

## Setup

### Docker-Compose

The repository has a `compose.yaml` file that contains the
configuration for docker-compose.

Note that it has some marked lines that should be removed or edited
in a production environment.

Other than that, simply run `docker-compose up -d` to launch the
containers.

That's it. The API should now be available under
`http://localhost:3005` and the Grafana UI should be accessible
from `http://localhost:3004`.

> **Note:** The data sources for Grafana have to be configured
> manually. For that, use the URLs `http://prometheus:9090` and
> `http://loki:3100`.

## Kubernetes

In order to run the application on Kubernetes, a running Kubernetes
cluster and the CLI tool Kubectl is required.

The file `deployment.yaml` serves as a manifest for Kubernetes and
has all the settings and configurations required for running the
application.

The configuration can be applied with
`kubectl apply -f deployment.yaml`.


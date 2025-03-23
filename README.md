# MediaDB

MediaDB is a RestAPI that serves metadata for media such as movies,
shows and books.

## Technologies

The API itself is a basic TypeScript application that uses
**ExpressJS** to serve the data and is built on **BunJS**.

The data is hosted on a **MongoDB** instance, while logs and metrics
are sent to a **Loki** and a **Prometheus** instance.

## Running/Building the Project

### Working Locally

Read the following files for information on how to run the project
locally:

- [MediaDB Server](./server/README.md)
- [MediaDB Client](./client/README.md)

### Building the Docker Image

With the following command, a Docker image can be built:

```bash
docker build -t mdb .
```

### Running the Development Environment

Using the compose configuration, a full development environment can
be spun up. For that, run following command:

```bash
docker-compose up -d --build
```

### Deploying on Kubernetes

The repository contains a `deployment.yaml`, that contains a complete
manifest for deployment on Kubernetes.

With the folling command, the configuration can be applied on a
running cluster:

```bash
kubectl apply -f deployment.yaml
```

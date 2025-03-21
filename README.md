# MediaDB

MediaDB is a RestAPI that serves metadata for media such as movies,
shows and books.

## Technologies

The API itself is a basic TypeScript application that uses
**ExpressJS** to serve the data.

The data is hosted on a **MongoDB** instance, while logs and metrics
are sent to a **Loki** and a **Prometheus** instance.

## Reference

See following resources for additional information:

- [Setup](./docs/setup.md)
- [API Reference](./docs/api-reference.md)

## API Structure

This image shows the general structure of the application with its
classes:

![mdb-classes.png](./docs/img/mdb-classes.png)

> **Note:** The database structure is not yet implemented like this.
> This is, however, the aim of this project.

## Database Structure

The following illustration shows the data structures from within the
MongoDB database:

![mdb-database_structure.png](./docs/img/mdb-database_structure.png)

> **Note:** The database structure is not yet implemented like this.
> This is, however, the aim of this project.

## System Design

![mdb-system_designs.png](./docs/img/mdb-system_designs.png)

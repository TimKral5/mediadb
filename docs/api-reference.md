# MediaDB API Reference

## `GET /api/v0/movies/:id`

Fetch detailed information about a movie.

### Parameters

- `:id` - The ID of the entry

## `GET /api/v0/movie-collections/:id` **(WIP)**

Fetch detailed information about a movie collection.

### Parameters

- `:id` - The ID of the entry

## `GET /api/v0/shows/:id`

Fetch detailed information about a show.

### Parameters

- `:id` - The ID of the entry

## `GET /api/v0/movies`

Search through the movies.

### Query Parameters

- `q` - The search query (parts of the title or description)

## `GET /api/v0/movie-collections` **(WIP)**

Search through the movie-collections.

### Query Parameters

- `q` - The search query (parts of the title or description)

## `GET /api/v0/shows`

Search through the shows.

### Query Parameters

- `q` - The search query (parts of the title or description)

## `POST /api/v0/movies` **(WIP)**

Create a new entry for movies.

### Body

- **Type:** JSON
- **Data:** Movie Data

### Query Parameters

- `q` - The search query (parts of the title or description)

## `POST /api/v0/movie-collections` **(WIP)**

Create a new entry for movie collections.

### Body

- **Type:** JSON
- **Data:** Movie Collection Data

## `POST /api/v0/shows` **(WIP)**

Create a new entry for shows.

### Body

- **Type:** JSON
- **Data:** Show Data

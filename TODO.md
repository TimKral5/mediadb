# MediaDB TODO

## Milestones
### 1. Read Access

#### API

- [x] Movie API
  - [x] Get with id
  - [x] Search by title or description
- [x] Show API
  - [x] Get with id
  - [x] Search by title or description
- [ ] MovieCollection API
  - [ ] Get with id
  - [ ] Search by title or description

#### Configuration

- [x] Environment Variables
  - [x] Read environment variable `NODE_ENV` for testing routine
  - [x] Read port from environment variable `MDB_PORT`
  - [x] Read mongodb url from environment variable `MDB_MONGODB_URL`
  - [x] Read loki database url from environment variable `MDB_LOKI_URL`

#### Testing

##### Testing Data

- [ ] Movies
- [ ] Shows
  - [ ] Seasons
  - [ ] Episodes
- [ ] MovieCollections

##### Testing Routine

- [ ] On development, apply testing routine
  - [ ] Load testing data to database
  - [ ] Execute tests
  - [ ] Cleanup database

##### Data Models

- [ ] Test `Movie` class
- [ ] Test `Show` class
- [ ] Test `MovieCollection` class

##### MVC Models

- [ ] Test function `MovieModel.getMovie`
- [ ] Test function `MovieModel.searchMovies`
- [ ] Test function `MovieModel.getCollection`
- [ ] Test function `MovieModel.searchCollections`
- [ ] Test function `ShowModel.getShow`
- [ ] Test function `ShowModel.searchShows`

##### Routes

- [ ] Test route `GET /movie?id={id}`
- [ ] Test route `GET /movie/search?q={query}`
- [ ] Test route `GET /movie/collection?id={id}`
- [ ] Test route `GET /movie/collection/search?q={query}`
- [ ] Test route `GET /show?id={id}`
- [ ] Test route `GET /show/search?q={query}`

### 2. CR Operations

#### API

- [ ] Movie API
  - [ ] Create from JSON
- [ ] Show API
  - [ ] Create from JSON
- [ ] MovieCollection API
  - [ ] Create from JSON

#### Testing

##### MVC Models

- [ ] Test function `MovieModel.createMovie`
- [ ] Test function `MovieModel.createMovieCollection`
- [ ] Test function `ShowModel.createShow`

##### Routes

- [ ] Test route `POST /movie` with JSON body
- [ ] Test route `POST /movie/collection` with JSON body
- [ ] Test route `POST /show` with JSON body

### 3. CRD Operations

//

### 4. CRUD Operations

//

### 5. API Keys

//


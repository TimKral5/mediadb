import {
  Movie,
  MovieCollection,
  Show,
  SemVer
} from 'mediadb-shared';

import {
  URL
} from 'url';

import { MissingEndpointError } from './errors/MissingEndpointError';

import { type SearchQuery } from './SearchQuery';

export class MediaDbConnection {
  constructor(
    private endpoint: string
  ) { }
  
  async testConnection(): Promise<boolean> {
    const res = await fetch(this.endpoint);
    return res.ok;
  }

  /**
   * Fetches the details of a movie with the specified ID
   * @param id The target ID
   */
  async getMovie(id: string): Promise<Movie> {
    const res = await fetch(`${this.endpoint}/movies/${id}`);

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = await res.json();
    return new Movie(data);
  }

  /**
   * Fetches the details of a movie collection with the specified ID
   * @param id The target ID
   */
  async getMovieCollection(id: string): Promise<MovieCollection> {
    const res = await fetch(`${this.endpoint}/movie-collections/${id}`);

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = await res.json();
    return new MovieCollection(data);
  }

  /**
   * Fetches the details of a show with the specified ID
   * @param id The target ID
   */
  async getShow(id: string): Promise<Show> {
    const res = await fetch(`${this.endpoint}/shows/${id}`, );

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = await res.json();
    return new Show(data);
  }

  /**
   * Searches through the movie entries and returns all matches
   * @param query The search parameters
   */
  async searchMovies(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/movies`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    
    const res = await fetch(url);

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = <Partial<Movie>[]>(await res.json());
    return data.map(movie => new Movie(movie));
  }

  /**
   * Searches through the movie collection entries and returns all
   * matches
   * @param query The search parameters
   */
  async searchMovieCollections(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/movie-collections`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    
    const res = await fetch(url);

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = <Partial<MovieCollection>[]>(await res.json());
    return data.map(movie => new MovieCollection(movie));
  }

  /**
   * Searches through the show entries and returns all matches
   * @param query The search parameters
   */
  async searchShows(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/shows`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    
    const res = await fetch(url);

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    const data = <Partial<Show>[]>(await res.json());
    return data.map(movie => new Show(movie));
  }

  /**
   * Creates a new movie from the specified data and returns the new
   * ID
   * @param data The data that is inserted
   */
  async createMovie(data: Movie): Promise<string> {
    const res = await fetch(`${this.endpoint}/movies`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).new_id;
  }

  /**
   * Creates a new movie collection from the specified data and
   * returns the new ID
   * @param data The data that is inserted
   */
  async createMovieCollection(data: MovieCollection): Promise<string> {
    const res = await fetch(`${this.endpoint}/movie-collections`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).new_id;
  }

  /**
   * Creates a new show from the specified data and returns the new
   * ID
   * @param data The data that is inserted
   */
  async createShow(data: Show): Promise<string> {
    const res = await fetch(`${this.endpoint}/shows`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).new_id;
  }

  /**
   * Updates a movie with the specified data and returns the ID
   * @param data The updated data
   */
  async updateMovie(data: Movie): Promise<string> {
    const res = await fetch(`${this.endpoint}/movies/${data.id}`, {
      body: JSON.stringify(data.dump()),
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).new_id;
  }

  /**
   * Updates a show with the specified data and returns the ID
   * @param data The updated data
   */
  async updateShow(data: Show): Promise<string> {
    const res = await fetch(`${this.endpoint}/shows/${data.id}`, {
      body: JSON.stringify(data.dump()),
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).new_id;
  }

  /**
   * Deletes a movie entry with the specified ID
   * @param id The target ID
   */
  async deleteMovie(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/movies/${id}`, {
      method: 'DELETE'
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).is_successful;
  }

  /**
   * Deletes a movie collection entry with the specified ID
   * @param id The target ID
   */
  async deleteMovieCollection(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/movie-collections/${id}`, {
      method: 'DELETE'
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).is_successful;
  }

  /**
   * Deletes a show entry with the specified ID
   * @param id The target ID
   */
  async deleteShow(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/shows/${id}`, {
      method: 'DELETE'
    });

    if (res.status === 404) {
      throw new MissingEndpointError(
        res.headers.get('X-API-Version') ?? 'v0.0.0');
    }

    return (await res.json()).is_successful;
  }
}

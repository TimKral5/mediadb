import {
  Movie,
  MovieCollection,
  Show
} from 'mediadb-shared';

import {
  URL
} from 'url';

import { type SearchQuery } from './SearchQuery';

export class MediaDbConnection {
  constructor(
    private endpoint: string
  ) { }
  
  async testConnection(): Promise<boolean> {
    const res = await fetch(this.endpoint);
    return res.ok;
  }

  async getMovie(id: string): Promise<Movie> {
    const res = await fetch(`${this.endpoint}/movies/${id}`);
    const data = await res.json();
    return new Movie(data);
  }

  async getMovieCollection(id: string): Promise<MovieCollection> {
    const res = await fetch(`${this.endpoint}/movie-collections/${id}`);
    const data = await res.json();
    return new MovieCollection(data);
  }

  async getShow(id: string): Promise<Show> {
    const res = await fetch(`${this.endpoint}/shows/${id}`, );
    const data = await res.json();
    return new Show(data);
  }

  async searchMovies(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/movies`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    console.log(url.toString());
    
    const res = await fetch(url);
    const data = <Partial<Movie>[]>(await res.json());
    return data.map(movie => new Movie(movie));
  }

  async searchMovieCollections(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/movie-collections`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    console.log(url.toString());
    
    const res = await fetch(url);
    const data = <Partial<MovieCollection>[]>(await res.json());
    return data.map(movie => new MovieCollection(movie));
  }

  async searchShows(query: Partial<SearchQuery>) {
    const url = new URL(`${this.endpoint}/shows`);

    Object.entries(query)
      .forEach(prop => url.searchParams
        .append(prop[0], <string>prop[1]));
    console.log(url.toString());
    
    const res = await fetch(url);
    const data = <Partial<Show>[]>(await res.json());
    return data.map(movie => new Show(movie));
  }

  async createMovie(data: Movie): Promise<string> {
    const res = await fetch(`${this.endpoint}/movies`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return (await res.json()).new_id;
  }
  
  async createMovieCollection(data: MovieCollection): Promise<string> {
    const res = await fetch(`${this.endpoint}/movie-collections`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return (await res.json()).new_id;
  }
  
  async createShow(data: Show): Promise<string> {
    const res = await fetch(`${this.endpoint}/shows`, {
      body: JSON.stringify(data.dump()),
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return (await res.json()).new_id;
  }

  async updateMovie(data: Movie): Promise<string> {
    const res = await fetch(`${this.endpoint}/movies/${data.id}`, {
      body: JSON.stringify(data.dump()),
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return (await res.json()).new_id;
  }

  async updateShow(data: Show): Promise<string> {
    const res = await fetch(`${this.endpoint}/shows/${data.id}`, {
      body: JSON.stringify(data.dump()),
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      }
    });
    return (await res.json()).new_id;
  }

  async deleteMovie(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/movies/${id}`, {
      method: 'DELETE'
    });
    return (await res.json()).is_successful;
  }

  async deleteMovieCollection(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/movie-collections/${id}`, {
      method: 'DELETE'
    });
    return (await res.json()).is_successful;
  }

  async deleteShow(id: string): Promise<boolean> {
    const res = await fetch(`${this.endpoint}/shows/${id}`, {
      method: 'DELETE'
    });
    return (await res.json()).is_successful;
  }
}

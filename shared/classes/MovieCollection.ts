
import { Media } from './Media';
import { Movie } from './Movie';

export class MovieCollection extends Media {
  public genres: string[];
  public movies: Movie[];

  constructor();
  constructor(data: Partial<MovieCollection>);
  
  constructor(data: Partial<MovieCollection> = {}) {
    super(data);
    let movies: { [key: string]: any }[] =
      (<{ [key: string]: any }>data)['_movies']
        ?? data['movies']
        ?? [];

    this.genres = data['genres'] ?? [];
    this.movies = movies.map(movie => new Movie(movie));
  }

  dump(keepId = false) {
    const obj = super.dump(keepId);
    obj.movies = this.movies.map(movie => movie.dump(true));
    return obj;
  }
}

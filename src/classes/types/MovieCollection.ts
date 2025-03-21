
import Media from './Media';
import Movie from './Movie';

export default class MovieCollection extends Media {
  public genres: string[];
  public movies: Movie[];

  constructor();
  constructor(data: Partial<MovieCollection>);
  
  constructor(data: Partial<MovieCollection> = {}) {
    super(data);
    let movies: { [key: string]: any }[] = (<{ [key: string]: any }>data)['_movies'] ?? data['movies'] ?? [];

    this.genres = data['genres'] ?? [];
    this.movies = movies.map(movie => new Movie(movie));
  }
}

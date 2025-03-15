
import Media from './Media';
import Movie from './Movie';

export default class MovieCollection extends Media {
  public genres: string[];
  public movies: Movie[];

  constructor();
  constructor(data: { [key: string]: any });
  
  constructor(data: { [key: string]: any } = {}) {
    super(data);
    const movies: { [key: string]: any }[] = data['movies'] ?? [];

    this.genres = data['genres'] ?? [];
    this.movies = movies.map(movie => new Movie(movie));
  }
}

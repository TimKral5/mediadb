import Media from './Media';

export default class Movie extends Media {
  public genres: string[];
  
  constructor();
  constructor(data: Partial<Movie>);

  constructor(data: Partial<Movie> = {}) {
    super(data);
    this.genres = data.genres ?? [];
  }
}

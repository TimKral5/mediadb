import Media from './Media';

export default class Movie extends Media {
  public genres: string[];
  
  constructor();
  constructor(data: { [key: string]: any });

  constructor(data: { [key: string]: any } = {}) {
    super(data);
    this.genres = data['genres'] ?? [];
  }
}

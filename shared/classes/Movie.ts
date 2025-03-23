import { Media } from './Media';

export class Movie extends Media {
  public genres: string[];
  
  constructor();
  constructor(data: Partial<Movie>);

  constructor(data: Partial<Movie> = {}) {
    super(data);
    this.genres = data.genres ?? [];
  }

  dump(keepId = false) {
    const obj = super.dump(keepId);
    return obj;
  }
}

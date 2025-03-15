
import Media from './Media';
import Season from './Season';

export default class Show extends Media {
  public genres: string[];
  public seasons: Season[];

  constructor();
  constructor(data: { [key: string]: any });
  
  constructor(data: { [key: string]: any } = {}) {
    super(data);
    const seasons: { [key: string]: any }[] = data['seasons'] ?? [];

    this.genres = data['genres'] ?? [];
    this.seasons = seasons.map(season => new Season(season));
  }
}

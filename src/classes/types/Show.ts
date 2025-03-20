
import Media from './Media';
import Season from './Season';

export default class Show extends Media {
  public genres: string[];
  public seasons: Season[];

  constructor();
  constructor(data: Partial<Show>);
  
  constructor(data: Partial<Show> = {}) {
    super(data);
    const seasons = data.seasons ?? [];

    this.genres = data.genres ?? [];
    this.seasons = seasons.map(season => new Season(season));
  }
}

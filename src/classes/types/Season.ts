
import Media from './Media';
import Episode from './Episode';

export default class Season extends Media {
  public episodes: Episode[];

  constructor();
  constructor(data: { [key: string]: any });
  
  constructor(data: { [key: string]: any } = {}) {
    super(data);
    const episodes: { [key:string]: any }[] = data['episodes'] ?? [];

    this.episodes = episodes.map(episode => new Episode(episode));
  }
}

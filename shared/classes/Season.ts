
import { Media } from './Media';
import { Episode } from './Episode';

export class Season extends Media {
  public episodes: Episode[];

  constructor();
  constructor(data: Partial<Season>);
  
  constructor(data: Partial<Season> = {}) {
    super(data);
    const episodes = data.episodes ?? [];

    this.episodes = episodes.map(episode => new Episode(episode));
  }

  dump(keepId = false) {
    const obj = super.dump(keepId);
    obj.epidodes = this.episodes.map(episode => episode.dump());
    return obj;
  }
}

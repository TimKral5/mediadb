
import Media from './Media';

export default class Episode extends Media {
  constructor();
  constructor(data: Partial<Episode>);
  
  constructor(data: Partial<Episode> = {}) {
    super(data);
  }
}

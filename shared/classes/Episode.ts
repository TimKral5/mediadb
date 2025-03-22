
import { Media } from './Media';

export class Episode extends Media {
  constructor();
  constructor(data: Partial<Episode>);
  
  constructor(data: Partial<Episode> = {}) {
    super(data);
  }

  dump() {
    const obj = super.dump();
    return obj;
  }
}

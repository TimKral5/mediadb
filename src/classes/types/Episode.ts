
import Media from './Media';

export default class Episode extends Media {
  constructor();
  constructor(data: { [key: string]: any });
  
  constructor(data: { [key: string]: any } = {}) {
    super(data);
  }
}

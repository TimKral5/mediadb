import { Translation } from './Translation';
import { Source } from './Source';

export abstract class Media {
  public id: string | undefined;
  public title: Translation[];
  public description: Translation[];
  public sources: Source[];

  constructor();
  constructor(data: Partial<Media>);

  constructor(data: Partial<Media> = {}) {
    const titles = data.title ?? [];
    const descriptions = data.description ?? [];
    const sources = data.sources ?? [];

    this.id = (<{ [key: string]: any }>data)['_id']
      ?? data.id ?? undefined;
    this.title = titles.map(title => new Translation(title));
    this.description = descriptions.map(
      title => new Translation(title)
    );
    this.sources = sources.map(source => new Source(source));
  }

  dump() {
    const obj: { [key: string]: any } = { ...this };
    obj.id = undefined;
    obj._id = undefined;
    return obj;
  }
}

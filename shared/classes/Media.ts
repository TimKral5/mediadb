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

  dump(keepId = false) {
    const obj: { [key: string]: any } = {};
    Object.entries(this).forEach(
      el => {
        const [key, value] = el;
        if (key === 'id' || key === '_id') {
          if (keepId)
            obj['_id'] = value;
          return;
        }
        obj[key] = value;
      });
    return obj;
  }
}

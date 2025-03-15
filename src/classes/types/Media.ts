import Translation from './Translation';
import Source from './Source';

export default class Media {
  public title: Translation[];
  public description: Translation[];
  public sources: Source[];

  constructor();
  constructor(data: { [key: string]: any });

  constructor(data: { [key: string]: any } = {}) {
    const titles: { [key: string]: any }[] = data['title'] ?? [];
    const descriptions: { [key: string]: any }[] = data['description'] ?? [];
    const sources: { [key: string]: any }[] = data['source'] ?? [];

    this.title = titles.map(title => new Translation(title));
    this.description = descriptions.map(title => new Translation(title));
    this.sources = sources.map(source => new Source(source));
  }
}

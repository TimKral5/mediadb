
export class Source {
  public type: string;
  public source: string;
  public data: string;

  constructor();
  constructor(data: { [key: string]: any });
  constructor(data: { [key: string]: any } = {}) {
    this.type = data['type'] ?? 'NULL';
    this.source = data['source'] ?? 'NULL';
    this.data = data['data'] ?? '';
  }
}

export default class Translation {
  public language: string;
  public text: string;

  constructor();
  constructor(data: { [key: string]: any });

  constructor(data: { [key: string]: any } = {}) {
    this.language = data['language'] ?? 'NULL';
    this.text = data['text'] ?? '';
  }
}

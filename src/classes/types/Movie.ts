import Genre from './Genre';

export default class Movie {
  constructor(
    public id: string,
    public title: string,
    public description: string,
    public genres: Genre[]) {
  }
}

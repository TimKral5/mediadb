import { config } from '../../config';
import { ObjectId } from 'mongodb';

import MongoUtils from '../utils/MongoUtils';
import Movie from '../types/Movie';
import MovieCollection from '../types/MovieCollection';
import MvcComponent from '../types/MvcComponent';
import type IModel from '../../interfaces/IModel';

export default class MovieModel
  extends MvcComponent
  implements IModel {
  
  createCollection(collName: string) {
    MongoUtils.initCollection(this.db, collName, coll => {
      coll.createIndex({
        'title.text': 'text',
        'description.text': 'text'
      });
    });
  }

  async getMovie(id: string): Promise<Movie | {}> {
    const collection = this.db.collection(config.tables['Movie']);
    const result = await collection
      .findOne({ _id: new ObjectId(id) });

    if (result)
      return new Movie(<any>result);
    return {};
  }

  async searchMovies(query: string): Promise<Movie[]> {
    const collection = this.db.collection(config.tables['Movie']);
    const results = await collection
      .find({ $text: { $search: query } })
      .toArray();

    this.logger.log(JSON.stringify(results));

    const arr = results.map(item => new Movie(<object>item));
    return arr;
  }

  async getCollection(id: string) {
    const collection = this.db.collection(config.tables['MovieCollection']);
    const result = await collection
      .findOne({ _id: new ObjectId(id) });

    if (result)
      return new Movie(<any>result);
    return {};
  }

  async searchCollections(query: string) {
    const collection = this.db.collection(config.tables['MovieCollection']);
    const results = await collection
      .find({ $text: { $search: query } })
      .toArray();

    this.logger.log(JSON.stringify(results));

    const arr = results.map(item => new MovieCollection(<object>item));
    return arr;
  }

  async createMovie(data: Partial<Movie>): Promise<string> {
    const movie = new Movie(data);
    return (await this.db
      .collection(config.tables['Movie'])
      .insertOne(movie)).insertedId.toString();
  }
}

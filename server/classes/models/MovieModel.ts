import { config } from '../../config';
import { ObjectId } from 'mongodb';

import {
  Movie,
  MovieCollection
} from 'mediadb-shared';

import MongoUtils from '../utils/MongoUtils';
import MvcComponent from '../types/MvcComponent';
import type IModel from '../../interfaces/IModel';

export default class MovieModel
  extends MvcComponent
  implements IModel {
  
  createCollections() {
    MongoUtils.initCollection(this.db, config.tables['Movie'], coll => {
      coll.createIndex({
        'title.text': 'text',
        'description.text': 'text'
      });
    });

    MongoUtils.initCollection(this.db, config.tables['MovieCollection'], coll => {
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

    const arr = results.map(item => new Movie(<object>item));
    return arr;
  }

  async getCollection(id: string) {
    const collection = this.db.collection(config.tables['MovieCollection']);
    const results = await collection
      .aggregate([
        { $match: { _id: new ObjectId(id) } },
        {
          $lookup: {
            from: 'mdb_movies',
            as: '_movies',
            localField: 'movies',
            foreignField: '_id'
          }
        }
      ]).toArray();

    if (results.length > 0)
      return new MovieCollection(<any>results[0]);
    return {};
  }

  async searchCollections(query: string) {
    const collection = this.db.collection(config.tables['MovieCollection']);
    const results = await collection
      .aggregate([
        {
          $match: {
            $text: {
              $search: query
            }
          }
        },
        {
          $lookup: {
            from: 'mdb_movies',
            as: '_movies',
            localField: 'movies',
            foreignField: '_id'
          }
        }
      ]).toArray();

    const arr = results.map(item => new MovieCollection(<object>item));
    return arr;
  }

  async createMovie(data: Partial<Movie>): Promise<ObjectId> {
    const movie = new Movie(data);

    const obj: { [key: string]: any } = { ...movie };
    obj.id = undefined;
    obj._id = undefined;
    
    return (await this.db
      .collection(config.tables['Movie'])
      .insertOne(obj)).insertedId;
  }

  async createMovieCollection(data: Partial<MovieCollection>): Promise<ObjectId> {
    const coll = new MovieCollection(data);

    const obj: { [key: string]: any } = { ...coll };
    obj.id = undefined;
    obj._id = undefined;

    const promises: Promise<ObjectId>[] = [];
    for (let i = 0; i < coll.movies.length; i++) {
      const movie = coll.movies[i];
      promises.push(this.createMovie(movie));
    }

    obj.movies = await Promise.all(promises);

    return (await this.db
      .collection(config.tables['MovieCollection'])
      .insertOne(obj)).insertedId;
  }
}

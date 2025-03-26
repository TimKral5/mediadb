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
 
  /**
   * Sets up all the collections related to this model
   */
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

  /**
   * Fetches the details of a movie from it's ID
   * @param id The id of the entry
   */
  async getMovie(id: string): Promise<Movie | {}> {
    const collection = this.db.collection(config.tables['Movie']);
    const result = await collection
      .findOne({ _id: new ObjectId(id) });

    if (result)
      return new Movie(<any>result);
    return {};
  }

  /**
   * Searches through all the movies with the given query and returns
   * the results
   * @param query The search query that is used in the search
   */
  async searchMovies(query: string): Promise<Movie[]> {
    const collection = this.db.collection(config.tables['Movie']);
    const results = await collection
      .find({ $text: { $search: query } })
      .toArray();

    const arr = results.map(item => new Movie(<object>item));
    return arr;
  }

  /**
   * Fetches the details of a movie collection from it's ID
   * @param id The id of the entry
   */
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

  /**
   * Searches through all the movie collections with the given query
   * and returns the results
   * @param query The search query that is used in the search
   */
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

  /**
   * Creates a new movie entry and returns the ID
   * @param data The movie data that is inserted
   */
  async createMovie(data: Partial<Movie>): Promise<ObjectId> {
    const movie = new Movie(data);

    const obj = movie.dump();
    return (await this.db
      .collection(config.tables['Movie'])
      .insertOne(obj)).insertedId;
  }

  /**
   * Creates a new movie collection entry and returns the ID
   * @param data The movie collection data that is inserted
   */
  async createMovieCollection(data: Partial<MovieCollection>): Promise<ObjectId> {
    const coll = new MovieCollection(data);

    const obj = coll.dump();
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

  /**
   * Updates the data for a movie from it's ID
   * @param id The ID of the movie entry that is updated
   * @param data The updated movie data
   */
  async updateMovie(id: string, data: Partial<Movie>) {
    const _id = new ObjectId(id);
    const movie = new Movie(data);

    return (await this.db.collection(config.tables['Movie'])
      .updateOne({ _id }, { $set: movie.dump() }))
      .upsertedId ?? _id;
  }

  /**
   * Deletes a movie entry from it's ID
   * @param id The ID of the movie entry that is deleted
   */
  async deleteMovie(id: string): Promise<boolean> {
    const _id = new ObjectId(id);

    return (await this.db
      .collection(config.tables['Movie'])
      .deleteOne({ _id })).deletedCount > 0;
  }

  /**
   * Deletes a movie collection entry from it's ID
   * @param id The ID of the movie colletion entry that is deleted
   */
  async deleteMovieCollection(id: string): Promise<boolean> {
    const _id = new ObjectId(id);

    return (await this.db
      .collection(config.tables['MovieCollection'])
      .deleteOne({ _id })).deletedCount > 0;
  }
}

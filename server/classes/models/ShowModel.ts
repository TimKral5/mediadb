import { config } from '../../config';
import { ObjectId } from 'mongodb';

import {
  Show
} from 'mediadb-shared';

import MongoUtils from '../utils/MongoUtils';
import MvcComponent from '../types/MvcComponent';
import type IModel from '../../interfaces/IModel';

export default class ShowModel
  extends MvcComponent
  implements IModel {
  
  createCollections() {
    MongoUtils.initCollection(this.db, 'mdb_shows', coll => {
      coll.createIndex({
        'title.text': 'text',
        'description.text': 'text'
      });
    });
  }

  async getShow(id: string): Promise<Show | {}> {
    const collection = this.db.collection(config.tables['Show']);
    const result = await collection
      .findOne({ _id: new ObjectId(id) });

    if (result)
      return new Show(<any>result);
    return {};
  }

  async searchShows(query: string): Promise<Show[]> {
    const collection = this.db.collection(config.tables['Show']);
    const results = await collection
      .find({ $text: { $search: query } })
      .toArray();

    const arr = results.map(item => new Show(<object>item));
    return arr;
  }

  async createShow(data: Partial<Show>): Promise<ObjectId> {
    const show = new Show(data);

    const obj = show.dump();
    return (await this.db
      .collection(config.tables['Show'])
      .insertOne(obj)).insertedId;
  }

  async updateShow(id: string, data: Partial<Show>): Promise<ObjectId> {
    const _id = new ObjectId(id);
    const show = new Show(data);

    const obj = show.dump();
    return (await this.db
      .collection(config.tables['Show'])
      .updateOne({ _id }, { $set: obj }))
      .upsertedId ?? _id;
  }

  async deleteShow(id: string): Promise<boolean> {
    const _id = new ObjectId(id);

    return (await this.db
      .collection(config.tables['Show'])
      .deleteOne({ _id })).deletedCount > 0;
  }
}

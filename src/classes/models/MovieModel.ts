
import { Db, Collection } from 'mongodb';

export class MovieModel {
  private db: Db;
  private collection: Collection | undefined;

  constructor(db: Db) {
    this.db = db;
  }

  async init(): Promise<Collection> {
    this.collection = await this.db.createCollection('mdb_movie');
    return this.collection;
  }

  async getById() {
    
  }
}

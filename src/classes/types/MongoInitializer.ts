import { MongoClient, Db } from 'mongodb';

export default class MongoInitializer {
  private db: Db;
  constructor(client: MongoClient, dbName: string) {
    this.db = client.db(dbName);
  }

  async createCollections(collections: string[]): Promise<MongoInitializer> {
    await Promise.all(
      collections.map(coll => this.db.createCollection(coll))
    );
    return this;
  }
}


import { Db, Collection } from 'mongodb';

export default class MongoUtils {
  static async initCollection(
    db: Db,
    collName: string,
    onCreate: (coll: Collection) => void
  ) {
    const arr = await db.listCollections({ name: collName })
      .toArray();

    if (arr.length === 0) {
      const coll = await db.createCollection(collName);
      onCreate(coll);
      return coll;
    }
  }
}

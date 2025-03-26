
import { Db, Collection } from 'mongodb';

export default class MongoUtils {
  /**
   * Checks if collection exists and creates it if it does not exist
   * @param db The database context
   * @param collName The name of the target collection
   * @param onCreate Callback function that is triggered if the
   * collection is created
   */
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

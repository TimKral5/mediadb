import { config } from '../../config';

import { ObjectId } from 'mongodb';

import Show from '../types/Show';
import MvcComponent from '../types/MvcComponent';

export default class ShowModel
  extends MvcComponent {

  getShowById(id: ObjectId): Show {
    throw new Error('Not implemented');
  }

  searchShow(query: string): Show[] {
    const collection = this.db.collection(config.tables['Show']);
    collection.find({  });
  }
}

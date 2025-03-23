import * as mediadb from '../index.ts';
(<any>globalThis)['mediadb'] = mediadb;
(<any>globalThis)['conn'] =
  new mediadb.MediaDbConnection('http://localhost:3005/api/v0');

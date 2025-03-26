import { test, expect } from 'bun:test';
import {
  MediaDbConnection,
  MovieCollection
} from 'mediadb-client';

const endpoint = process.env['MDB_ENDPOINT'];
if (!endpoint)
  throw new Error('Missing Environment Variable');

const conn = new MediaDbConnection(endpoint);

const coll = new MovieCollection({
  title: [
    {
      language: 'en',
      text: 'A Movie Collection'
    }
  ]
});

let collId = await conn.createMovieCollection(coll);
coll.id = collId;

test('Create MovieCollection', () => {
  expect(collId).toBeDefined();
});

test('Get MovieCollection', async () => {
  const data = await conn.getMovieCollection(collId);
  expect(data).toBeDefined();
});

test('Search MovieCollections', async () => {
  const data = await conn.searchMovieCollections({
    q: 'collection'
  });
  expect(data.length).toBeGreaterThan(0);
});

test('Delete MovieCollection', async () => {
  coll.movies.forEach(movie => conn.deleteMovie(movie.id ?? ''));
  const res = await conn.deleteMovieCollection(collId);
  expect(res).toBeTrue();
});

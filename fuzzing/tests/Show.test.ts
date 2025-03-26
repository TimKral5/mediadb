import { test, expect } from 'bun:test';
import {
  MediaDbConnection,
  Show
} from 'mediadb-client';

const endpoint = process.env['MDB_ENDPOINT'];
if (!endpoint)
  throw new Error('Missing Environment Variable');

const conn = new MediaDbConnection(endpoint);

const show = new Show({
  title: [
    {
      language: 'en',
      text: 'A Show'
    }
  ]
});

let showId = await conn.createShow(show);
show.id = showId;

test('Create Show', () => {
  expect(showId).toBeDefined();
});

test('Get Show', async () => {
  const data = await conn.getShow(showId);
  expect(data).toBeDefined();
});

test('Search Shows', async () => {
  const data = await conn.searchShows({
    q: 'show'
  });
  expect(data.length).toBeGreaterThan(0);
});

test('Update Show', async () => {
  const desc = 'Added Description';
  show.description = [{
    language: 'en',
    text: desc
  }];
  const newId = await conn.updateShow(show);
  const data = await conn.getShow(newId);
  
  expect(data.description[0].text).toBe(desc);
});

test('Delete Show', async () => {
  const res = await conn.deleteShow(showId);
  expect(res).toBeTrue();
});

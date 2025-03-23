import { test, expect } from 'bun:test';
import {
  MediaDbConnection,
  Movie
} from 'mediadb-client';

const endpoint = process.env['MDB_ENDPOINT'];
if (!endpoint)
  throw new Error('Missing Environment Variable');

const conn = new MediaDbConnection(endpoint);

const movie = new Movie({
  title: [
    {
      language: 'en',
      text: 'A Movie'
    }
  ]
});

let movieId = await conn.createMovie(movie);
movie.id = movieId;

test('Create Movie', () => {
  expect(movieId).toBeDefined();
});

test('Get Movie', async () => {
  const data = await conn.getMovie(movieId);
  expect(data).toBeDefined();
});

test('Search Movies', async () => {
  const data = await conn.searchMovies({
    q: 'movie'
  });
  expect(data.length).toBeGreaterThan(0);
});

test('Update Movie', async () => {
  const desc = 'Added Description';
  movie.description = [{
    language: 'en',
    text: desc
  }];
  const newId = await conn.updateMovie(movie);
  const data = await conn.getMovie(newId);
  
  expect(data.description[0].text).toBe(desc);
});

test('Delete Movie', async () => {
  const res = await conn.deleteMovie(movieId);
  expect(res).toBeTrue();
});

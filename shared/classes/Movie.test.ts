import { test, describe, beforeEach, expect } from 'bun:test';
import { Movie } from './Movie.ts';

const data: Partial<Movie> = {
  title: [],
  description: [],
  genres: [],
  sources: []
};

let movie: Movie | undefined ;
beforeEach(() => movie = new Movie(data));

describe('Movie Properties', () => {
  test('title', () => {
    expect((<Movie>movie).title).toBeArray();
  });
  test('description', () => {
    expect((<Movie>movie).description).toBeArray();
  });
  test('sources', () => {
    expect((<Movie>movie).sources).toBeArray();
  });
  test('genres', () => {
    expect((<Movie>movie).genres).toBeArray();
  });
  test('id', () => {
    expect((<Movie>movie).id).toBeUndefined();
  });
});

test.todo('Complete testing');


import {
  MediaDbConnection,
  Movie
} from './index.ts';

const conn = new MediaDbConnection('http://localhost:3005/api/v0');

const id = await conn.createMovie(new Movie({
  title: [
    {
      language: 'en',
      text: 'A New Show'
    }
  ],
}));

const movie = await conn.getMovie(id);
console.log(movie);
movie.title.push({
  language: 'de',
  text: 'Eine Neue Serie'
});

conn.updateMovie(movie);

const newMovie = await conn.getMovie(id);
console.log(newMovie);

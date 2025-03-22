
import {
  MediaDbConnection,
  Show,
  Season
} from './index.ts';

const conn = new MediaDbConnection('http://localhost:3005/api/v0');

const id = await conn.createShow(new Show({
  title: [
    {
      language: 'en',
      text: 'A New Show'
    }
  ],
  seasons: [
    new Season({
      title: [
        {
          language: 'en',
          text: 'Season 1'
        }
      ]
    })
  ]
}));

const show = await conn.getShow(id);
console.log(show);


import type {
  Express,
  Request,
  Response
} from 'express';

import { Registry } from 'prom-client';
import { Db } from 'mongodb';

import type IController from '../../interfaces/IController';

import type Logger from '../types/Logger';
import MvcComponent from '../types/MvcComponent';
import MovieModel from '../models/MovieModel';

export default class MovieController
  extends MvcComponent
  implements IController {

  private model: MovieModel;

  constructor(
    logger: Logger,
    registry: Registry,
    db: Db
  ) {
    super(logger, registry, db);
    this.model = this.initMvcComponent(MovieModel);
    this.model.createCollection('mdb_movies');
  }

  private async getMovie(req: Request, res: Response) {
    const id = <string>req.params['id'];

    try {
      const data = await this.model.getMovie(id);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async searchMovies(req: Request, res: Response) {
    const query = <string>req.query['q'];

    try {
      const data = await this.model.searchMovies(query);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async getCollection(req: Request, res: Response) {
    const id = <string>req.params['id'];

    try {
      const data = await this.model.getCollection(id);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async searchCollections(req: Request, res: Response) {
    const query = <string>req.query['q'];

    try {
      const data = await this.model.searchCollections(query);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  registerRoutes(baseRoute: string, app: Express) {
    app.get(`${baseRoute}/movies/:id`, this.getMovie.bind(this));
    app.get(`${baseRoute}/movies`, this.searchMovies.bind(this));
    app.get(`${baseRoute}/movie-collections/:id`, this.getCollection.bind(this));
    app.get(`${baseRoute}/movie-collections`, this.searchCollections.bind(this));
  }
}

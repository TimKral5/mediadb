
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
    this.model.createCollections();
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

  private async createMovie(req: Request, res: Response) {
    try {
      let data = {};
      try {
        data = req.body;
        this.logger.log(JSON.stringify(data));
      }
      catch (_err) {
        const err = <Error>_err;
        res.status(400).json({});
        this.logger.error(err.toString());
        return;
      }
      const id = await this.model.createMovie(data);
      res.json({
        new_id: id.toString()
      });
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async createMovieCollection(req: Request, res: Response) {
    try {
      let data = {};
      try {
        data = req.body;
        this.logger.log(JSON.stringify(data));
      }
      catch (_err) {
        const err = <Error>_err;
        res.status(400).json({});
        this.logger.error(err.toString());
        return;
      }
      const id = await this.model.createMovieCollection(data);
      res.json({
        new_id: id.toString()
      });
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
    app.post(`${baseRoute}/movies`, this.createMovie.bind(this));
    
    app.get(`${baseRoute}/movie-collections/:id`, this.getCollection.bind(this));
    app.get(`${baseRoute}/movie-collections`, this.searchCollections.bind(this));
    app.post(`${baseRoute}/movie-collections`, this.createMovieCollection.bind(this));
  }
}

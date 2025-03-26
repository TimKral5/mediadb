
import type {
  Express,
  Request,
  Response
} from 'express';

import {
  Registry,
  Counter
} from 'prom-client';

import { Db } from 'mongodb';

import type IController from '../../interfaces/IController';

import type Logger from '../types/Logger';
import MvcComponent from '../types/MvcComponent';
import MovieModel from '../models/MovieModel';

export default class MovieController
  extends MvcComponent
  implements IController {

  private model: MovieModel;

  private counters: { [key: string]: Counter };

  /**
   * Adds a new counter to the counters object.
   * @param name Name of the counter
   * @param help The helptext that is displayed in the metrics route
   */
  private createCounter(name: string, help: string) {
    this.counters[name] = new Counter({
      name: `mdb_${name}_calls`,
      help
    });
  }

  /**
   * Creates and registers all counters used in the routes.
   */
  private registerCounters() {
    const helptext = 'Amount of calls to Endpoint';

    this.createCounter('get_movie', helptext);
    this.createCounter('search_movies', helptext);
    this.createCounter('create_movie', helptext);
    this.createCounter('update_movie', helptext);
    this.createCounter('delete_movie', helptext);

    this.createCounter('get_movie_collection', helptext);
    this.createCounter('search_movie_collections', helptext);
    this.createCounter('create_movie_collection', helptext);
    this.createCounter('update_movie_collection', helptext);
    this.createCounter('delete_movie_collection', helptext);

    Object.values(this.counters).forEach(
      counter => this.registry.registerMetric(counter));
  }

  constructor(
    logger: Logger,
    registry: Registry,
    db: Db
  ) {
    super(logger, registry, db);
    this.model = this.initMvcComponent(MovieModel);
    this.model.createCollections();

    this.counters = {};
    this.registerCounters();
  }

  private async getMovie(req: Request, res: Response) {
    this.logger.debug(`GET ${req.url}`);
    this.counters['get_movie'].inc();

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
    this.logger.debug(`GET ${req.url}`);
    this.counters['search_movies'].inc();

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
    this.logger.debug(`GET ${req.url}`);
    this.counters['get_movie_collection'].inc();

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
    this.logger.debug(`GET ${req.url}`);
    this.counters['search_movie_collections'].inc();

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
    this.logger.debug(`POST ${req.url}`);
    this.counters['create_movie'].inc();

    try {
      let data = {};
      try {
        data = req.body;
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
    this.logger.debug(`POST ${req.url}`);
    this.counters['create_movie_collection'].inc();

    try {
      let data = {};
      try {
        data = req.body;
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

  private async updateMovie(req: Request, res: Response) {
    this.logger.debug(`PUT ${req.url}`);
    this.counters['update_movie'].inc();

    try {
      const id = req.params['id'];
      let data = {};
      try {
        data = req.body;
      }
      catch (_err) {
        const err = <Error>_err;
        res.status(400).json({});
        this.logger.error(err.toString());
        return;
      }
      const _id = await this.model.updateMovie(id, data);
      res.json({
        new_id: _id.toString()
      });
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async deleteMovie(req: Request, res: Response) {
    this.logger.debug(`DELETE ${req.url}`);
    this.counters['delete_movie'].inc();

    try {
      const id = req.params['id'];
      const isSuccessful = await this.model.deleteMovie(id);
      res.json({
        is_successful: isSuccessful
      });
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async deleteMovieCollection(req: Request, res: Response) {
    this.logger.debug(`DELETE ${req.url}`);
    this.counters['delete_movie_collection'].inc();

    try {
      const id = req.params['id'];
      const isSuccessful = await this.model.deleteMovieCollection(id);
      res.json({
        is_successful: isSuccessful
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
    app.put(`${baseRoute}/movies/:id`, this.updateMovie.bind(this));
    app.delete(
      `${baseRoute}/movies/:id`,
      this.deleteMovie.bind(this)
    );
    
    app.get(
      `${baseRoute}/movie-collections/:id`,
      this.getCollection.bind(this)
    );
    app.get(
      `${baseRoute}/movie-collections`,
      this.searchCollections.bind(this)
    );
    app.post(
      `${baseRoute}/movie-collections`,
      this.createMovieCollection.bind(this)
    );
    app.delete(
      `${baseRoute}/movie-collections/:id`,
      this.deleteMovieCollection.bind(this)
    );
  }
}

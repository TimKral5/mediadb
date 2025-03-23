
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
import ShowModel from '../models/ShowModel';

export default class ShowController
  extends MvcComponent
  implements IController {

  private model: ShowModel;

  private counters: { [key: string]: Counter };

  private createCounter(name: string, help: string) {
    this.counters[name] = new Counter({
      name: `mdb_${name}_calls`,
      help
    });
  }

  private registerCounters() {
    const helptext = 'Amount of calls to Endpoint';

    this.createCounter('get_show', helptext);
    this.createCounter('search_shows', helptext);
    this.createCounter('create_show', helptext);
    this.createCounter('update_show', helptext);

    Object.values(this.counters).forEach(
      counter => this.registry.registerMetric(counter));
  }

  constructor(
    logger: Logger,
    registry: Registry,
    db: Db
  ) {
    super(logger, registry, db);
    this.model = new ShowModel(logger, registry, db);
    this.model.createCollections();

    this.counters = {};
    this.registerCounters();
  }

  private async getShow(req: Request, res: Response) {
    this.logger.debug(`GET ${req.url}`);
    this.counters['get_show'].inc();
    try {
      const id = <string>req.params['id'];
      const data = await this.model.getShow(id);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async searchShows(req: Request, res: Response) {
    this.logger.debug(`GET ${req.url}`);
    this.counters['search_shows'].inc();
    const query = <string>req.query['q'];

    try {
      const data = await this.model.searchShows(query);
      res.json(data);
    }
    catch (_err) {
      const err = <Error>_err;
      this.logger.error(err.toString());
      res.status(500).end();
    }
  }

  private async createShow(req: Request, res: Response) {
    this.logger.debug(`POST ${req.url}`);
    this.counters['create_show'].inc();
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
      const id = await this.model.createShow(data);
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

  private async updateShow(req: Request, res: Response) {
    this.logger.debug(`PUT ${req.url}`);
    this.counters['update_show'].inc();
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
      const _id = await this.model.updateShow(id, data);
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

  registerRoutes(baseRoute: string, app: Express) {
    app.get(`${baseRoute}/shows/:id`, this.getShow.bind(this));
    app.get(`${baseRoute}/shows`, this.searchShows.bind(this));
    app.post(`${baseRoute}/shows`, this.createShow.bind(this));
    app.put(`${baseRoute}/shows/:id`, this.updateShow.bind(this));
  }
}

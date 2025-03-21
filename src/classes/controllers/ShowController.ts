
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
import ShowModel from '../models/ShowModel';

export default class ShowController
  extends MvcComponent
  implements IController {

  private model: ShowModel;

  constructor(
    logger: Logger,
    registry: Registry,
    db: Db
  ) {
    super(logger, registry, db);
    this.model = new ShowModel(logger, registry, db);
    this.model.createCollection('mdb_shows');
  }

  private async getShow(req: Request, res: Response) {
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

  registerRoutes(baseRoute: string, app: Express) {
    app.get(`${baseRoute}/show/:id`, this.getShow.bind(this));
    app.get(`${baseRoute}/shows`, this.searchShows.bind(this));
  }
}

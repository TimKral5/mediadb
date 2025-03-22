
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
    this.model.createCollections();
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

  private async createShow(req: Request, res: Response) {
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
      const id = await this.model.createShow(data);
      res.json({
        new_id: id
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
  }
}

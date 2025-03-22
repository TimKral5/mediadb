import type { Express, Request, Response } from 'express';

import type IController from '../../interfaces/IController';
import type IModel from '../../interfaces/IModel';
import MvcComponent from '../types/MvcComponent';

export default class PrometheusApiController
  extends MvcComponent
  implements IController {

  private async getMetrics(_: Request, res: Response) {
    res.set('Content-Type', this.registry.contentType);
    res.end(await this.registry.metrics());
  }

  registerRoutes(baseRoute: string, app: Express): void {
    app.get(baseRoute, this.getMetrics.bind(this));
  }
}

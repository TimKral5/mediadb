import type { Express, Request, Response } from 'express';
import { Registry } from 'prom-client';

import type IController from '../../interfaces/IController';
import MvcComponent from '../types/MvcComponent';

export default class PrometheusApiController
  extends MvcComponent
  implements IController {
  private async getMetrics(_: Request, res: Response) {
    try {
      res.set('Content-Type', this.registry.contentType);
      res.end(await this.registry.metrics());
    }
    catch (err) {
      this.logger.error((<Error>err).toString());
    }
  }

  registerRoutes(baseRoute: string, app: Express): void {
    app.get(baseRoute, this.getMetrics.bind(this));
  }
}

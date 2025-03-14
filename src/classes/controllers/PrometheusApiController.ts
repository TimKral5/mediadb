import type { Express, Request, Response } from 'express';
import { Registry } from 'prom-client';

import type IController from '../../interfaces/IController';
import RequestHandler from '../types/RequestHandler';
import Logger from '../types/Logger';

export default class PrometheusApiController implements IController {
  private baseRoute: string;
  private registry: Registry;
  private logger: Logger;

  constructor(baseRoute: string, registry: Registry, logger: Logger) {
    this.baseRoute = baseRoute;
    this.registry = registry;
    this.logger = logger;
  }

  private async getMetrics(req: Request, res: Response) {
    try {
      res.set('Content-Type', this.registry.contentType);
      res.end(await this.registry.metrics());
    }
    catch (err) {
      this.logger.error((<Error>err).toString());
    }
  }

  registerRoutes(app: Express): void {
    app.get(
      this.baseRoute,
      this.getMetrics.bind(this)
    );
  }
}

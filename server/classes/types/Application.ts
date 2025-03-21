
import type { Express } from 'express';
import express from 'express';
import { Registry } from 'prom-client';
import { Db } from 'mongodb';

import type IController from '../../interfaces/IController';
import Logger from './Logger';

type MvcType<T> = new (logger: Logger, registry: Registry, db: Db) => T;

export default abstract class Application {
  public express: Express;

  constructor(
    private logger: Logger,
    private registry: Registry,
    private db: Db) {
    this.express = express();
    this.express.use(express.json());
    this.express.use(express.urlencoded({ extended: true }));
  }

  public initMvcCompontent<T>(_class: MvcType<T>) {
    return new _class(
      this.logger,
      this.registry,
      this.db
    );
  }

  public registerControllers(controllers: Array<[string, IController]>) {
    for (let i = 0; i < controllers.length; i++) {
      const [ baseRoute, controller ] = controllers[i];
      controller.registerRoutes(baseRoute, this.express);
    }
  }

  public launch(port: number) {
    this.express.listen(port, () => {
      this.logger.log(`Listening on port ${port}...`);
    });
  }
}

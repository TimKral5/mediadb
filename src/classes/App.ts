
import type { Express } from 'express';
import express from 'express';
import type IController from '../interfaces/IController';
import Logger from './types/Logger';

export default class App {
  public express: Express;
  private logger: Logger;

  constructor(logger: Logger) {
    this.express = express();
    this.logger = logger;
  }

  public registerControllers(controllers: Array<IController>) {
    for (let i = 0; i < controllers.length; i++) {
      const controller = controllers[i];
      controller.registerRoutes(this.express);
    }
  }

  public launch(port: number) {
    this.express.listen(port, () => {
      this.logger.log(`Listening on port ${port}...`);
    });
  }
}

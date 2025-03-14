import type { Request, Response } from 'express';
import Logger from './Logger';

export default class RequestHandler {
  requestHandler: (req: Request, res: Response) => void;
  logger: Logger;

  constructor(handler: (req: Request, res: Response) => void, logger: Logger) {
    this.requestHandler = handler;
    this.logger = logger;
  }

  private handleRequest(req: Request, res: Response) {
    try {
      this.requestHandler(req, res);
    }
    catch (err) {
      this.logger.error((<Error>err).toString());
    }
  }

  getHandler() {
    return this.handleRequest.bind(this);
  }
}

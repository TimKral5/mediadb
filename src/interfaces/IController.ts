
import type { Express } from 'express';

export default interface IController {
  registerRoutes(baseRoute: string, app: Express): void;
}


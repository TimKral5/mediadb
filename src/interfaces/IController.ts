
import type { Express } from 'express';

export default interface IController {
  registerRoutes(app: Express): void;
}


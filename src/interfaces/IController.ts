
import type { Express } from 'express';
import type IModel from '../interfaces/IModel';

export default interface IController {
  registerRoutes(baseRoute: string, app: Express): void;
}


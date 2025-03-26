
import type { Express } from 'express';

export default interface IController {
  /**
   * Registers all routes with the specified base route
   * @param baseRoute The base route
   * @param app The express app instance
   */
  registerRoutes(baseRoute: string, app: Express): void;
}


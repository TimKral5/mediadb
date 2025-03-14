import type IController from '../../interfaces/IController';
import type { Express } from 'express';

export default class AuthController implements IController {
  constructor(private baseRoute: string) {

  }
  registerRoutes(app: Express) {
    app.get(`${this.baseRoute == '/' ? '' : this.baseRoute}/api_key`);
  }
}

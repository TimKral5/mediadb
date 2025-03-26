
import Application from './types/Application';

import PrometheusApiController from './controllers/PrometheusApiController';
import ShowController from './controllers/ShowController';
import MovieController from './controllers/MovieController';

export default class App extends Application {
  main() {
    this.express.use('/', (_, res, next) => {
      res.setHeader('X-API-Version', 'v0.2.0');
      next();
    });
    this.registerControllers([
      ['/api/v0', this.initMvcCompontent(ShowController)],
      ['/metrics', this.initMvcCompontent(PrometheusApiController)],
      ['/api/v0', this.initMvcCompontent(MovieController)]
    ]);
  }
}

import { MongoClient } from 'mongodb';
import { Registry } from 'prom-client';
import { transports, format } from 'winston';
import LokiTransport from 'winston-loki';

import App from './classes/App';
import PrometheusApiController from './classes/controllers/PrometheusApiController';
import Logger from './classes/types/Logger';
import MongoInitializer from './classes/types/MongoInitializer';

// MongoDB
const MONGODB_URL = 'mongodb://root:example@mongodb:27017';
const client = await MongoClient.connect(MONGODB_URL);
await new MongoInitializer(client, 'mdb').createCollections([
  'mdb_movies',
  'mdb_shows',
  'mdb_books'
]);

// Prometheus
const registry = new Registry();

// Loki
const logger = new Logger([
  new LokiTransport({
    host: 'http://loki:3100',
    labels: { app: 'mdb' }
  }),
  new transports.Console({
    format: format.combine(format.simple(), format.colorize())
  })
]);

const app = new App(logger);

app.registerControllers([
  new PrometheusApiController('/metrics', registry, logger)
]);

app.launch(3000);

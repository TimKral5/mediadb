import { MongoClient } from 'mongodb';
import { Registry } from 'prom-client';
import { transports, format } from 'winston';
import LokiTransport from 'winston-loki';

import App from './classes/App';
import Logger from './classes/types/Logger';

// MongoDB
const MONGODB_URL = 'mongodb://root:example@mongodb:27017';
const client = await MongoClient.connect(MONGODB_URL);
const db = client.db('mdb');

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

const app = new App(logger, registry, db);

app.main();
app.launch(3000);

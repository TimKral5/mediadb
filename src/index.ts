import { MongoClient } from 'mongodb';
import { Registry } from 'prom-client';
import { transports, format } from 'winston';
import LokiTransport from 'winston-loki';

import App from './classes/App';
import Logger from './classes/types/Logger';

const ENV = (process.env['NODE_ENV'] ?? 'development').toLowerCase();
const MDB_PORT = parseInt(process.env['MDB_PORT'] ?? '3000');
const MDB_MONGODB_URL = process.env['MDB_MONGODB_URL'];
const MDB_LOKI_URL = process.env['MDB_LOKI_URL'];

if (!MDB_MONGODB_URL || !MDB_LOKI_URL)
  throw new Error('ERROR: Missing environment variables.');

if (ENV === 'development') {
  // testing routine goes here
  
}

// MongoDB
const client = await MongoClient.connect(MDB_MONGODB_URL);
const db = client.db('mdb');

// Prometheus
const registry = new Registry();

// Loki
const logger = new Logger([
  new LokiTransport({
    host: MDB_LOKI_URL,
    labels: { app: 'mdb' }
  }),
  new transports.Console({
    format: format.combine(format.simple(), format.colorize())
  })
]);

const app = new App(logger, registry, db);

app.main();
app.launch(MDB_PORT);

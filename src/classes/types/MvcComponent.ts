
import Logger from './Logger';
import { Registry } from 'prom-client';
import { Db } from 'mongodb';

export default abstract class MvcComponent {
  constructor(
    readonly logger: Logger,
    readonly registry: Registry,
    readonly db: Db
  ) {}
}

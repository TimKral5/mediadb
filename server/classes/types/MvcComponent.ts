
import Logger from './Logger';
import { Registry } from 'prom-client';
import { Db } from 'mongodb';

type MvcConstructor<T> =
  new (logger: Logger, registry: Registry, db: Db) => T;

export default abstract class MvcComponent {
  constructor(
    readonly logger: Logger,
    readonly registry: Registry,
    readonly db: Db
  ) {}

  initMvcComponent<T>(_class: MvcConstructor<T>) {
    return new _class(
      this.logger,
      this.registry,
      this.db
    );
  }
}

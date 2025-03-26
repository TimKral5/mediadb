
import Logger from './Logger';
import {
  Registry
} from 'prom-client';
import { Db } from 'mongodb';

type MvcType<T> =
  new (logger: Logger, registry: Registry, db: Db) => T;

export default abstract class MvcComponent {
  constructor(
    readonly logger: Logger,
    readonly registry: Registry,
    readonly db: Db,
  ) {}

  /**
   * A utility function for the initialization of a MVC component
   * @param _class The type that is initialized
   */
  initMvcComponent<T>(_class: MvcType<T>) {
    return new _class(
      this.logger,
      this.registry,
      this.db
    );
  }
}

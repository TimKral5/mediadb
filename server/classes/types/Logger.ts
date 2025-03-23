
import { createLogger, Logger as WinstonLogger } from 'winston';

export default class Logger {
  private logger: WinstonLogger;

  constructor(transports: Array<any>) {
    this.logger = createLogger({
      transports: transports,
      levels: {
        emerg: 0,
        alert: 1,
        crit: 2,
        error: 3,
        warning: 4,
        notice: 5,
        info: 6,
        debug: 7
      },
      level: 'debug'
    });
  }

  log(message: string) {
    this.logger.info(message);
  }

  warn(message: string) {
    this.logger.warn(message);
  }

  error(message: string) {
    this.logger.error(message);
  }

  debug(message: string) {
    this.logger.debug(message);
  }
}

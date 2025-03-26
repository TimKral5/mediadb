
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

  /**
   * Write an info message into the logs
   * @param message The message that is logged
   */
  log(message: string) {
    this.logger.info(message);
  }

  /**
   * Write a warning into the logs
   * @param message The message that is logged
   */
  warn(message: string) {
    this.logger.warn(message);
  }

  /**
   * Write an error message into the logs
   * @param message The message that is logged
   */
  error(message: string) {
    this.logger.error(message);
  }

  /**
   * Write debug info into the logs
   * @param message The message that is logged
   */
  debug(message: string) {
    this.logger.debug(message);
  }
}

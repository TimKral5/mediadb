
import { createLogger, Logger as WinstonLogger } from 'winston';

export default class Logger {
  private logger: WinstonLogger;

  constructor(transports: Array<any>) {
    this.logger = createLogger({
      transports: transports
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
}

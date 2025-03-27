import { SemVer, SemVerCompatibility } from 'mediadb-shared';

import { ErrorContext } from './ErrorContext';

import { MissingEndpointError } from './errors/MissingEndpointError';
import { MalformedRequestError } from './errors/MalformedRequestError';
import { ServerSideError } from './errors/ServerSideError';

export class ErrorHandler {
  constructor(
    private context: Partial<ErrorContext>
  ) {}

  private compareVersions(): SemVerCompatibility {
    const apiVersion = new SemVer(this.context.apiVersion ?? 'v0.0.0');
    const compatibility = SemVer.currentVersion.compareTo(apiVersion);
    return compatibility;
  }

  private detectCause(): string {
    if (this.compareVersions() !== SemVerCompatibility.COMPATIBLE)
      return 'Incompatible API versions';
    return 'Unknown cause';
  }

  handleStatus(status: number) {
    if (status === 200)
      return;

    const cause = this.detectCause();

    if (status === 404)
      throw new MissingEndpointError(cause);
    if (status === 400)
      throw new MalformedRequestError(cause);
    if (status === 500)
      throw new ServerSideError(cause);
  }
}

import { SemVer, SemVerCompatibility } from 'mediadb-shared';

export class MissingEndpointError extends Error {
  static readonly clientApiVersion = new SemVer('v0.2.0');
  clientVersion: string;
  serverVersion: string;
  versionCompatibility: SemVerCompatibility;

  constructor(apiVersion: string);
  constructor(apiVersion: SemVer);

  constructor(apiVersion: string | SemVer) {
    let serverVersion: SemVer;
    if (typeof apiVersion === 'string')
      serverVersion = new SemVer(apiVersion);
    else
      serverVersion = apiVersion;

    const clientVersion = MissingEndpointError.clientApiVersion;

    super();
    this.name = 'MissingEndpointError';
    this.message = 'The targeted API endpoint does not exist';

    this.clientVersion = clientVersion.toString();
    this.serverVersion = serverVersion.toString();

    const compatibility = clientVersion.compareTo(serverVersion);
    this.versionCompatibility = compatibility;

    if (compatibility !== SemVerCompatibility.COMPATIBLE)
      this.cause = 'Incompatible API versions';
    else
      this.cause = 'Unknown cause';
  }
}

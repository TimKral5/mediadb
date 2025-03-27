export enum SemVerCompatibility {
  COMPATIBLE = 'compatible',
  DIFFERENT_MAJOR = 'different_major',
  NEWER_MINOR = 'newer_minor'
}

export class SemVer {

  public static currentVersion = new SemVer('v0.3.0');

  public major: number;
  public minor: number;
  public patch: number;

  constructor();
  constructor(version: string);

  constructor(version: string | undefined = undefined) {
    if (version === undefined) {
      this.major = 0;
      this.minor = 0;
      this.patch = 0;
      return;
    }

    const versionNumbers = version.substring(1).split('.');
    this.major = parseInt(versionNumbers[0]);
    this.minor = parseInt(versionNumbers[1]);
    this.patch = parseInt(versionNumbers[2]);
  }

  toString() {
    return `v${this.major}.${this.minor}.${this.patch}`;
  }

  /**
   * Compares this version to the provided one and checks, if they
   * are compatible.
   * @param version Reference version
   */
  compareTo(version: SemVer): SemVerCompatibility {
    if (this.major !== version.major)
      return SemVerCompatibility.DIFFERENT_MAJOR;
    if (this.minor > version.minor)
      return SemVerCompatibility.NEWER_MINOR;
    return SemVerCompatibility.COMPATIBLE;
  }
}

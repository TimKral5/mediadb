export class ServerSideError extends Error {
  constructor(
    public cause: string
  ) {
    super();
  }
}

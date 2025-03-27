export class MalformedRequestError extends Error {
  constructor(
    public cause: string
  ) {
    super();
  }
}

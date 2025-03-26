
export default interface IModel {
  /**
   * A function that handles the initialization of database
   * collections
   */
  createCollections(): void;
}

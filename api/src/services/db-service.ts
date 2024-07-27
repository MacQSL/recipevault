import { Repository } from "../repositories/repository";
import { Connection } from "../utils/database";

/**
 * Database Service class.
 *
 * @description Used to implement service classes which require a repository dependency.
 *
 * @class DBService
 * @property {Repository} repository - Repository dependency
 */
export class DBService {
  repository: Repository;

  constructor(connection: Connection) {
    this.repository = new Repository(connection);
  }
}

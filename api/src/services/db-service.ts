import { Repository } from '../repositories/repository.js';
import { Connection } from '../utils/database.js';

/**
 * Database Service class.
 *
 * Note: All database service classes should `implemement` this class.
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

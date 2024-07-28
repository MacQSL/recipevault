import { Connection } from '../utils/database.js';

/**
 * Base repository class.
 *
 * Note: All repository classes should `extend` this class.
 *
 * @class Repository
 * @property {Connection} connection - Knex transaction client
 */
export class Repository {
  connection: Connection;

  constructor(connection: Connection) {
    this.connection = connection;
  }
}

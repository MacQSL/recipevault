import { Connection } from "../utils/database";

/**
 * Base repository class.
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

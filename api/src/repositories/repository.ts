import { Connection } from "../utils/database";

export class Repository {
  connection: Connection;

  constructor(connection: Connection) {
    this.connection = connection;
  }
}

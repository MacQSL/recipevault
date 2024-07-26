import knex, { Knex } from "knex";

class Repository {
  knex: Knex;

  constructor() {
    this.knex = knex({
      client: "pg",
      connection: {
        host: process.env.DB_HOST,
        port: Number(process.env.DB_PORT),
        user: process.env.DB_USER_API,
        password: process.env.DB_USER_API_PASSWORD,
        database: process.env.DB,
      },
    });
  }
}

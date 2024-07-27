import knex from "knex";
import { Sql } from "sql-template-tag";
import { APIError500 } from "./error";

/**
 * Connection timeout milliseconds (5 seconds).
 *
 */
const CONNECTION_TIMEOUT_MS = 5000;

/**
 * Knex client config.
 *
 */
const _knex = knex({
  client: "pg",
  connection: {
    host: process.env.DB_HOST,
    port: Number(process.env.DB_PORT),
    user: process.env.DB_USER_API,
    password: process.env.DB_USER_API_PASSWORD,
    database: process.env.DB_DATABASE,
  },
  pool: { min: 0, max: 10 },
});

/**
 * Knex transaction client - singleton instance
 *
 */
const knexTransactionClient = _knex.transactionProvider();

/**
 * Database connection - Knex transaction provider instance.
 *
 */
export type Connection = {
  knex: knex.Knex.Transaction;
  sql(query: Sql): Promise<any>;
  commit(): void;
  rollback(): void;
};

/**
 * Get database connection instance, uses knex transactionProvider.
 *
 * Note: Calling this function will start the connection.
 * Connection will timeout after 5 seconds if transaction not commited / rolledback.
 *
 * @async
 * @returns {Promise<Connection>} Knex transaction provider instance
 */
export const getDBConnection = async (): Promise<Connection> => {
  const knexClient = await knexTransactionClient();

  // If connection opened and not closed after 5 seconds throw error
  setTimeout(() => {
    if (!knexClient.isCompleted()) {
      throw new APIError500("Transaction opened without being closed.");
    }
  }, CONNECTION_TIMEOUT_MS);

  return {
    knex: knexClient,
    sql: async (query: Sql) => knexClient.raw(query.sql, query.values),
    commit: () => {
      knexClient.commit();
    },
    rollback: () => {
      knexClient.rollback();
    },
  };
};

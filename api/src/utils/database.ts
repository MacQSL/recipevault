import knex from 'knex';
import { Sql } from 'sql-template-tag';
import { APIError500 } from './error.js';

/**
 * Connection timeout milliseconds (5 seconds).
 * Note: This could be placed into ENV (I think here is fine for now).
 *
 */
const CONNECTION_TIMEOUT_MS = 5000;

/**
 * Knex client config.
 *
 */
const _knex = knex({
  client: 'pg',
  connection: {
    host: process.env.DB_HOST,
    port: Number(process.env.DB_PORT),
    database: process.env.DB_DATABASE,
    user: process.env.DB_USER_API,
    password: process.env.DB_USER_API_PASSWORD
  },
  pool: { min: 0, max: 10 }
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
  /**
   * Execute query with knex query builder.
   *
   */
  knex: knex.Knex.Transaction;
  /**
   * Execute query with raw SQL (sql-template-tag).
   *
   */
  sql(query: Sql): Promise<any>;
  /**
   * Commit the transaction.
   *
   */
  commit(): void;
  /**
   * Rollback the transaction.
   *
   */
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
      throw new APIError500('Transaction opened without being closed.');
    }
  }, CONNECTION_TIMEOUT_MS);

  /**
   * Exectute the SQL-template-tag query with the knex transaction client.
   *
   * @async
   * @param {Sql} query - SQL statement
   * @returns {Promise<any[]>}
   */
  const sqlQuery = async (query: Sql): Promise<any[]> => {
    const response = await knexClient.raw(query.sql, query.values);
    return response.rows;
  };

  return {
    knex: knexClient,
    sql: sqlQuery,
    commit: knexClient.commit,
    rollback: knexClient.rollback
  };
};

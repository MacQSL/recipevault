import knex from "knex";

const CONNECTION_TIMEOUT_MS = 5000;

/**
 * Knex transaction provider instance.
 *
 */
type Connection = knex.Knex.Transaction;

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
    database: process.env.DB,
  },
  pool: { min: 0, max: 10 },
});

// Singleton instance
const knexTransactionClient = _knex.transactionProvider();

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
  const connection = await knexTransactionClient();

  setTimeout(() => {
    if (!connection.isCompleted()) {
      console.warn("TRANSACTION OPENED WITHOUT BEING CLOSED!");
    }
  }, CONNECTION_TIMEOUT_MS);

  return connection;
};

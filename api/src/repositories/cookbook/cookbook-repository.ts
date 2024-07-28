import { Repository } from '../repository.js';
import sql from 'sql-template-tag';

/**
 * Cookbook Repository class.
 *
 * @class CookbookRepository
 * @extends Repository
 */
export class CookbookRepository extends Repository {
  /**
   * Get cookbooks by user id.
   *
   * @async
   * @param {number} userId
   * @returns {*}
   */
  async getCookbooksByUserId(userId: number) {
    const sqlStatement = sql`
      SELECT
        c.cookbook_id,
        c.name,
        c.description
      FROM cookbook c
      LEFT JOIN user_cookbook u
      ON u.user_id = ${userId};
    `;

    return this.connection.sql(sqlStatement);
  }
}

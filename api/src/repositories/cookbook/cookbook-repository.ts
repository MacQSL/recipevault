import sql from 'sql-template-tag';
import { Repository } from '../repository.js';

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
      INNER JOIN user_cookbook u
      ON c.cookbook_id = u.cookbook_id
      WHERE u.user_id = ${userId};
    `;

    return this.connection.sql(sqlStatement);
  }
}

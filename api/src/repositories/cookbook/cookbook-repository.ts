import sql from 'sql-template-tag';
import { APIError404 } from '../../utils/error.js';
import { Repository } from '../repository.js';
import { ICookbookRecord } from './cookbook-repository.interface.js';

/**
 * Cookbook Repository class.
 *
 * @class CookbookRepository
 * @extends Repository
 */
export class CookbookRepository extends Repository {
  /**
   * Get cookbook by id.
   *
   * @async
   * @param {number} cookbookId
   * @returns {ICookbookRecord[]}
   */
  async getCookbookById(cookbookId: number): Promise<ICookbookRecord> {
    const sqlStatement = sql`
      SELECT
        c.cookbook_id,
        c.name,
        c.description
      FROM cookbook c
      WHERE c.cookbook_id = ${cookbookId};
    `;

    const response = await this.connection.sql(sqlStatement);

    if (!response.length) {
      throw new APIError404('Cookbook not found.', ['CookbookRepository->getCookbookById']);
    }

    return response[0];
  }

  /**
   * Get cookbooks by user id.
   *
   * @async
   * @param {number} userId
   * @returns {ICookbookRecord[]}
   */
  async getCookbooksByUserId(userId: number): Promise<ICookbookRecord[]> {
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

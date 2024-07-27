import { Repository } from "../repository";
import sql from "sql-template-tag";

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
        cookbook_id,
        name,
        description
      FROM cookbook
      JOIN user_cookbook u
      ON u.user_id = ${userId};
    `;

    return this.connection.sql(sqlStatement);
  }
}

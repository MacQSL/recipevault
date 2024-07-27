import { Repository } from "../repository";

export class CookbookRepository extends Repository {
  private cookbookColumns = [];

  /**
   * Get cookbooks by user id.
   *
   * @async
   * @param {number} userId
   * @returns {*}
   */
  async getCookbooksByUserId(userId: number) {
    return this.connection
      .select(this.cookbookColumns)
      .from("recipe")
      .where({ user_id: userId });
  }
}

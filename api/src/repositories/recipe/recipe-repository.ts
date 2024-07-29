import SQL from 'sql-template-tag';
import { APIError404 } from '../../utils/error.js';
import { Repository } from '../repository.js';

const test = 'test';

/**
 * Recipe Repository class.
 *
 * @class RecipeRepository
 * @extends Repository
 */
export class RecipeRepository extends Repository {
  /**
   * Get recipe by id.
   *
   * @async
   * @param {number} recipeId
   * @throws {APIError404} - Recipe not found
   * @returns {IRecipeRecord}
   */
  async getRecipeById(recipeId: number): Promise<IRecipeRecord> {
    const sqlStatement = SQL`
      SELECT
        recipe_id,
        cookbook_id,
        name,
        url,
        description
      FROM recipe
      WHERE recipe_id = ${recipeId};
    `;

    const response = await this.connection.sql(sqlStatement);

    if (!response.length) {
      throw new APIError404('Recipe not found.', ['RecipeRepository->getRecipeById']);
    }

    return response[0];
  }

  /**
   * Get recipes by cookbook id.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipeRecord[]>}
   */
  async getRecipesByCookbookId(cookbookId: number): Promise<IRecipeRecord[]> {
    const sqlStatement = SQL`
      SELECT
        recipe_id,
        cookbook_id,
        name,
        url,
        description
      FROM recipe
      WHERE cookbook_id = ${cookbookId};
    `;

    return this.connection.sql(sqlStatement);
  }
}

import { IRecipeRecord } from '../../repositories/recipe/recipe-repository.interface.js';
import { RecipeRepository } from '../../repositories/recipe/recipe-repository.js';
import { Connection } from '../../utils/database.js';
import { DBService } from '../db-service.js';

/**
 * Recipe Service class.
 *
 * @class RecipeService
 * @property {RecipeRepository} repository - Recipe repository dependency
 */
export class RecipeService implements DBService {
  repository: RecipeRepository;

  constructor(connection: Connection) {
    this.repository = new RecipeRepository(connection);
  }

  /**
   * Get recipe.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipeRecord>}
   */
  async getRecipe(cookbookId: number): Promise<IRecipeRecord> {
    // TODO: This should also get the ingredients
    return this.repository.getRecipeById(cookbookId);
  }

  /**
   * Get cookbook recipes.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipeRecord[]>}
   */
  async getCookbookRecipes(cookbookId: number): Promise<IRecipeRecord[]> {
    return this.repository.getRecipesByCookbookId(cookbookId);
  }
}

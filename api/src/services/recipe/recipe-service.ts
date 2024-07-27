import { RecipeRepository } from "../../repositories/recipe/recipe-repository";
import { Connection } from "../../utils/database";
import { DBService } from "../db-service";

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
   * @returns {Promise<IRecipe>}
   */
  async getRecipe(cookbookId: number): Promise<IRecipe> {
    return this.repository.getRecipeById(cookbookId);
  }

  /**
   * Get cookbook recipes.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipe[]>}
   */
  async getCookbookRecipes(cookbookId: number): Promise<IRecipe[]> {
    return this.repository.getRecipesByCookbookId(cookbookId);
  }
}

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
   * Get cookbook recipes.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipe[]>}
   */
  async getCookbookRecipes(cookbookId: number): Promise<IRecipe[]> {
    return this.repository.getRecipesFromCookbookId(cookbookId);
  }
}

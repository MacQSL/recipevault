import { APIService } from "./api-service";
import { IRecipe } from "./recipe-service.interface";

/**
 * @exports
 * @class RecipeService
 *
 */
export class RecipeService extends APIService {
  /**
   * Get recipe by id.
   *
   * @async
   * @returns {Promise<IRecipe>}
   */
  async getRecipe(cookbookId: number, recipeId: number): Promise<IRecipe> {
    const response = await this.client.get(
      `/cookbooks/${cookbookId}/recipes/${recipeId}`,
    );
    return response.data;
  }
}

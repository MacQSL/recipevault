import { ICookbookRecipes } from "./cookbook-service.interface";
import { APIService } from "./api-service";

/**
 * @exports
 * @class CookbookService
 *
 */
export class CookbookService extends APIService {
  /**
   * Get user cookbooks with recipes.
   *
   * @async
   * @returns {Promise<ICookbookRecipes[]>}
   */
  async getCookbooksWithRecipes(): Promise<ICookbookRecipes[]> {
    const response = await this.client.get("/cookbooks");
    return response.data;
  }
}

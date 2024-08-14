import { AxiosInstance } from "axios";

import { ICookbookRecipes } from "./cookbook-service.interface";

/**
 * @exports
 * @class CookbookService
 *
 */
export class CookbookService {
  api: AxiosInstance;

  constructor(api: AxiosInstance) {
    this.api = api;
  }

  /**
   * Get user cookbooks with recipes.
   *
   * @async
   * @returns {Promise<ICookbookRecipes>}
   */
  async getCookbooksWithRecipes(): Promise<ICookbookRecipes> {
    const response = await this.api.get("/cookbooks");
    console.log(response.data.data.data);
    return response.data.data;
  }
}

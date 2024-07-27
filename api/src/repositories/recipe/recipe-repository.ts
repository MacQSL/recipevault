import { Repository } from "../repository";

export class RecipeRepository extends Repository {
  private recipeColumns = [
    "recipe_id",
    "cookbook_id",
    "name",
    "url",
    "description",
  ];

  /**
   * Get recipe by id.
   *
   * @async
   * @param {number} recipeId
   * @throws {APIError404} - Recipe not found
   * @returns {IRecipe}
   */
  async getRecipeById(recipeId: number): Promise<IRecipe> {
    const response = await this.connection
      .select(this.recipeColumns)
      .from("recipe")
      .where({ recipe_id: recipeId });

    if (!response.length) {
      throw new APIError404("Recipe not found.", [
        "RecipeRepository->getRecipeById",
      ]);
    }

    return response[0];
  }

  /**
   * Get recipes by cookbook id.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipe[]>}
   */
  async getRecipesByCookbookId(cookbookId: number): Promise<IRecipe[]> {
    return this.connection
      .select(this.recipeColumns)
      .from("recipe")
      .where({ cookbook_id: cookbookId });
  }
}

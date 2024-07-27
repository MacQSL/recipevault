import { Repository } from "../repository";

export class RecipeRepository extends Repository {
  // omitting audit columns
  private recipeColumns = [
    "recipe_id",
    "cookbook_id",
    "name",
    "url",
    "description",
  ];

  /**
   * Get recipes from a cookbook id.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<IRecipe[]>}
   */
  async getRecipesFromCookbookId(cookbookId: number): Promise<IRecipe[]> {
    return this.connection
      .select(this.recipeColumns)
      .from("recipe")
      .where({ cookbook_id: cookbookId });
  }
}

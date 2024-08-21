import { useQuery } from "@tanstack/react-query";
import { client } from "../axios";
import { RecipeService } from "../service/recipe-service";

/**
 * Recipe query keys
 *
 */
export const RECIPE_QK = "recipe";

const recipeService = new RecipeService(client);

/**
 * Get recipe by id
 *
 */
export const useRecipeQuery = (cookbookId: number, recipeId: number) => {
  return useQuery({
    queryKey: [RECIPE_QK, cookbookId, recipeId],
    queryFn: async () => {
      return recipeService.getRecipe(cookbookId, recipeId);
    },
  });
};

import { queryOptions, useQuery } from "@tanstack/react-query";
import { client } from "../axios";
import { RecipeService } from "../../services/recipe-service";

const recipeService = new RecipeService(client);

export const recipeQueryOpts = (cookbookId: number, recipeId: number) => {
  return queryOptions({
    queryKey: ["recipe", cookbookId, recipeId],
    queryFn: async () => {
      return recipeService.getRecipe(cookbookId, recipeId);
    },
  });
};

/**
 * Get recipe by id
 *
 */
export const useRecipeQuery = (
  ...params: Parameters<typeof recipeQueryOpts>
) => {
  return useQuery(recipeQueryOpts(...params));
};

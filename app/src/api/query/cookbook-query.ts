import { queryOptions } from "@tanstack/react-query";
import { instance } from "../axios";
import { CookbookService } from "../service/cookbook-service";

// Instantiate cookbook service
const cookbookService = new CookbookService(instance);

/**

 * Get user cookbooks with recipes
 *
 * @returns {ICookbookRecipes}
 */
export const getCookbooksWithRecipes = () => {
  return queryOptions({
    queryKey: ["getCookbooksWithRecipes"],
    queryFn: async () => {
      cookbookService.getCookbooksWithRecipes();
    },
  });
};

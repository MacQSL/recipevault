import { useQuery } from "@tanstack/react-query";
import { client } from "../axios";
import { CookbookService } from "../service/cookbook-service";

/**
 * Cookbook query keys
 *
 */
export const COOKBOOKS_QK = "cookbooks";

const cookbookService = new CookbookService(client);

/**
 * Get user cookbooks with recipes
 *
 */
export const useCookbooksQuery = () => {
  return useQuery({
    queryKey: [COOKBOOKS_QK],
    queryFn: async () => {
      return cookbookService.getCookbooksWithRecipes();
    },
  });
};

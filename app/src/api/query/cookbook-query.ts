import { queryOptions, useQuery } from "@tanstack/react-query";
import { client } from "../axios";
import { CookbookService } from "../../services/cookbook-service";

const cookbookService = new CookbookService(client);

export const cookbooksQueryOpts = () => {
  return queryOptions({
    queryKey: ["cookbooks"],
    queryFn: async () => {
      return cookbookService.getCookbooksWithRecipes();
    },
  });
};

/**
 * Get user cookbooks with recipes
 *
 */
export const useCookbooksQuery = () => {
  return useQuery(cookbooksQueryOpts());
};

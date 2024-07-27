import { Connection } from "../../utils/database";
import { RecipeService } from "../recipe/recipe-service";

/**
 * Cookbook Service class.
 *
 * @class CookbookService
 * @property {RecipeService} recipeService  - Recipe service dependency
 */
export class CookbookService {
  recipeService: RecipeService;

  constructor(connection: Connection) {
    this.recipeService = new RecipeService(connection);
  }
}

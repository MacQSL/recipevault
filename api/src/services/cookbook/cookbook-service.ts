import { CookbookRepository } from "src/repositories/cookbook/cookbook-repository";
import { Connection } from "../../utils/database";
import { DBService } from "../db-service";
import { RecipeService } from "../recipe/recipe-service";

/**
 * Cookbook Service class.
 *
 * @class CookbookService
 * @property {CookbookRepository} repository  - Cookbook repository dependency
 * @property {RecipeService} recipeService  - Recipe service dependency
 */
export class CookbookService implements DBService {
  repository: CookbookRepository;
  recipeService: RecipeService;

  constructor(connection: Connection) {
    this.repository = new CookbookRepository(connection);
    this.recipeService = new RecipeService(connection);
  }
}

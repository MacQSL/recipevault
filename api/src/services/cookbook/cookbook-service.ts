import { CookbookRepository } from '../../repositories/cookbook/cookbook-repository.js';
import { Connection } from '../../utils/database.js';
import { DBService } from '../db-service.js';
import { RecipeService } from '../recipe/recipe-service.js';

/**
 * Cookbook Service class.
 *
 * @class CookbookService
 * @property {CookbookRepository} repository  - Cookbook repository dependency
 * @property {RecipeService} recipeService  - Recipe service dependency
 */
export class CookbookService implements DBService {
  repository: CookbookRepository;
  private recipeService: RecipeService;

  constructor(connection: Connection) {
    this.repository = new CookbookRepository(connection);
    this.recipeService = new RecipeService(connection);
  }

  async getUserCookbooks(userId: number) {
    return this.repository.getCookbooksByUserId(userId);
  }
}

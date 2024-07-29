import { ICookbookRecord, ICookbookWithRecipes } from '../../repositories/cookbook/cookbook-repository.interface.js';
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

  /**
   * Get a cookbook with recipes.
   *
   * Note: Recipes do not contain ingredients.
   *
   * @async
   * @param {number} cookbookId
   * @returns {Promise<ICookbookWithRecipes>} Cookbook and stub recipes
   */
  async getCookbookWithRecipes(cookbookId: number): Promise<ICookbookWithRecipes> {
    const [cookbook, recipes] = await Promise.all([
      this.repository.getCookbookById(cookbookId),
      this.recipeService.getCookbookRecipes(cookbookId)
    ]);

    return { ...cookbook, recipes };
  }

  /**
   * Get all user cookbooks.
   *
   * @async
   * @param {number} userId
   * @returns {Promise<ICookbookRecord>}
   */
  async getUserCookbooks(userId: number): Promise<ICookbookRecord[]> {
    return this.repository.getCookbooksByUserId(userId);
  }
}

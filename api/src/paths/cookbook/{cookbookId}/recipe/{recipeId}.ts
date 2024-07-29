import { RequestHandler } from 'express';
import { Operation } from 'express-openapi';
import { RecipeService } from '../../../../services/recipe/recipe-service.js';
import { getDBConnection } from '../../../../utils/database.js';

export const GET: Operation = [getRecipe()];

GET.apiDoc = {
  description: 'Get a recipe by id.',
  tags: ['recipe'],
  operationId: 'getRecipe',
  parameters: [
    {
      in: 'path',
      name: 'cookbookId',
      required: true
    },
    {
      in: 'path',
      name: 'recipeId',
      required: true
    }
  ],
  responses: {
    200: {
      description: 'Success',
      content: {
        'application/json': {
          schema: {
            type: 'object',
            required: ['recipe_id', 'name', 'url', 'description'],
            additionalProperties: false,
            properties: {
              recipe_id: {
                type: 'string'
              },
              name: {
                type: 'string'
              },
              description: {
                type: 'string',
                nullable: true
              }
            }
          }
        }
      }
    }
  }
};

export function getRecipe(): RequestHandler {
  return async (_req, res) => {
    const connection = await getDBConnection();
    const recipeService = new RecipeService(connection);

    try {
      const recipe = await recipeService.getRecipe(1);

      connection.commit();

      return res.status(200).json(recipe);
    } catch (err) {
      connection.rollback();
    }
  };
}

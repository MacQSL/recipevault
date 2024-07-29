import { RequestHandler } from 'express';
import { Operation } from 'express-openapi';
import { CookbookService } from '../../services/cookbook/cookbook-service.js';
import { getDBConnection } from '../../utils/database.js';

export const GET: Operation = [getUserCookbooks()];

GET.apiDoc = {
  description: 'Get a users cookbooks.',
  tags: ['cookbook'],
  operationId: 'getCookbook',
  responses: {
    200: {
      description: 'Success',
      content: {
        'application/json': {
          schema: {
            type: 'object',
            required: ['cookbook_id', 'name', 'description'],
            additionalProperties: false,
            properties: {
              cookbook_id: {
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

export function getUserCookbooks(): RequestHandler {
  return async (_req, res) => {
    const connection = await getDBConnection();
    const cookbookService = new CookbookService(connection);

    try {
      const cookbooks = await cookbookService.getUserCookbooks(2);

      connection.commit();

      return res.status(200).json(cookbooks);
    } catch (err) {
      connection.rollback();
    }
  };
}

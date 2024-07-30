import { OpenAPIV3 } from 'openapi-types';
import { errorSchema } from './schemas/error.js';

const apiDoc: OpenAPIV3.Document = {
  openapi: '3.0.0',
  info: {
    title: 'RecipeHub API.',
    version: '1.0.0'
  },
  components: {
    schemas: {
      Error: errorSchema
    },
    responses: {
      default: {
        description: 'Error response',
        content: {
          'application/json': {
            schema: {
              $ref: '#/components/schemas/Error'
            }
          }
        }
      }
    }
  },
  paths: {}
};

export default apiDoc;

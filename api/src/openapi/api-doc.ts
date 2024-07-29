import { OpenAPIV3 } from 'openapi-types';

const apiDoc: OpenAPIV3.Document = {
  openapi: '3.0.0',
  info: {
    title: 'RecipeHub API.',
    version: '1.0.0'
  },
  components: {
    schemas: {
      Error: {
        description: 'Error response object',
        required: ['name', 'status', 'message'],
        type: 'object',
        properties: {
          name: {
            type: 'string'
          },
          status: {
            description: 'HTTP response code',
            type: 'number'
          },
          message: {
            type: 'string'
          },
          errors: {
            type: 'array',
            items: {}
          }
        }
      }
    },
    responses: {
      default: {
        $ref: '#/components/schemas/Error'
      }
    }
  },
  paths: {}
};

export default apiDoc;

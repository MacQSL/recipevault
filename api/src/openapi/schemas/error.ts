import { OpenAPIV3 } from 'openapi-types';

export const errorSchema: OpenAPIV3.SchemaObject = {
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
};

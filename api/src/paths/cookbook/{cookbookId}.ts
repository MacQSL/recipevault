import { Request, Response } from 'express';
import { Operation } from 'express-openapi';

export const parameters = [
  {
    in: 'path',
    name: 'id',
    require: true,
    type: 'integer'
  }
];

export const GET: Operation = (_req: Request, res: Response) => {
  res.status(200).json('hi');
};

GET.apiDoc = {
  description: 'Get a cookbook.',
  tags: ['cookbook'],
  operationId: 'getCookbook'
};

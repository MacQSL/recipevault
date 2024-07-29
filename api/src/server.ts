import express from 'express';
import { initialize } from 'express-openapi';
import apiDoc from './openapi/api-doc.js';

const PORT = Number(process.env.API_PORT);

const app = express();

initialize({
  app,
  apiDoc: apiDoc,
  paths: './build/src/paths',
  routesGlob: '**/*.{ts,js}',
  routesIndexFileRegExp: /(?:index)?\.[tj]s$/,
  pathsIgnore: new RegExp('.(spec|test)$')
});

app.listen(PORT, async () => {
  console.log(`🍔 Server started on port ${PORT}`);
});

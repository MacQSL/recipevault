import express, { Request, Response } from 'express';

const PORT = Number(process.env.API_PORT);

const app = express();

app.get('/health', async (_req: Request, res: Response) => {
  res.send('🍔');
});

app.listen(PORT, async () => {
  console.log(`🍔 Server started on port ${PORT}`);
});

import express, { Request, Response } from "express";

const PORT = Number(process.env.API_PORT) || 3000;

const app = express();

app.get("/health", (_req: Request, res: Response) => {
  res.send("RecipeHub healthy :)");
});

app.listen(PORT, () => {
  console.log(`Server started on port ${PORT}`);
});

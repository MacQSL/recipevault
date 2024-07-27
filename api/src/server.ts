import express, { Request, Response } from "express";
import { getDBConnection } from "./utils/database";
import { RecipeRepository } from "./repositories/recipe/recipe-repository";

const PORT = Number(process.env.API_PORT) || 3000;

const app = express();

app.get("/health", async (_req: Request, res: Response) => {
  res.send("RecipeHub healthy :)");
});

app.listen(PORT, async () => {
  console.log(`Server started on port ${PORT}`);
  const connection = await getDBConnection();
  const recipeRepo = new RecipeRepository(connection);

  const test = await recipeRepo.getRecipesFromCookbookId(1);
  console.log(test);
});

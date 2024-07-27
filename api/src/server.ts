import express, { Request, Response } from "express";
import { getDBConnection } from "./utils/database";
import { CookbookService } from "./services/cookbook/cookbook-service";

const PORT = Number(process.env.API_PORT);

const app = express();

app.get("/health", async (_req: Request, res: Response) => {
  res.send("🍔");
});

app.listen(PORT, async () => {
  console.log(`🍔 Server started on port ${PORT}`);

  const connection = await getDBConnection();
  const service = new CookbookService(connection);
  const data = await service.getUserCookbooks(2);
  console.log(data);
});

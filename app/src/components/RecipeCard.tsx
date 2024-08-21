import { Avatar, Box, Card, Typography } from "@mui/joy";
import { IRecipe } from "../api/service/recipe-service.interface";

interface IRecipeCardProps {
  recipe: IRecipe;
}

export const RecipeCard = (props: IRecipeCardProps) => {
  return (
    <Card sx={{ width: 200 }}>
      <Typography level="title-lg">{props.recipe.name}</Typography>
      <Typography level="body-sm">{props.recipe.description}</Typography>
      <Box sx={{ display: "flex", gap: 1.5, mt: "auto" }}>
        <Avatar variant="soft" color="neutral">
          MD
        </Avatar>
        <div>
          <Typography level="body-xs">Created by</Typography>
          <Typography level="body-sm">Mac Deluca</Typography>
        </div>
      </Box>
    </Card>
  );
};

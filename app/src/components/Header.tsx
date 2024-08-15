import { Input, Sheet, Typography } from "@mui/joy";
import { ToggleMode } from "./ToggleMode";

export const Header = () => {
  return (
    <Sheet
      variant="outlined"
      sx={{
        display: "flex",
        justifyContent: "space-between",
        p: 2,
      }}
    >
      <Typography component="h1" level="h3">
        RecipeVault
      </Typography>
      <Input placeholder="Search for recipes..."></Input>
      <ToggleMode />
    </Sheet>
  );
};

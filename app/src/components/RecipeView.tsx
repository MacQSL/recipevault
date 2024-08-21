import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Input,
  Typography,
} from "@mui/joy";
import { SheetBox } from "./SheetBox";

export const RecipeView = () => {
  return (
    <SheetBox
      sx={{
        height: "100%",
        width: "100%",
        display: "flex",
        flexDirection: "column",
        gap: 2,
      }}
    >
      <Box>
        <Box sx={{ display: "flex" }}>
          <Typography level="h4" component="h3">
            Recipe Name
          </Typography>
          <Button sx={{ ml: "auto" }} size="sm">
            Edit
          </Button>
        </Box>
        <Typography level="body-md">Recipe description</Typography>
      </Box>
      <FormControl>
        <FormLabel>Ingredients</FormLabel>
        <Input placeholder="1 cup flour..." />
        <FormLabel>Steps</FormLabel>
        <Input />
      </FormControl>
    </SheetBox>
  );
};

import {
  Avatar,
  Box,
  Button,
  List,
  ListDivider,
  ListItem,
  Typography,
} from "@mui/joy";
import { SheetBox } from "./SheetBox";
import AddIcon from "@mui/icons-material/Add";

export const RecipeList = () => {
  return (
    <SheetBox
      sx={{
        height: "100%",
        width: 500,
      }}
    >
      <List>
        <Box sx={{ display: "flex" }}>
          <Typography level="h4" component="h2">
            Recipes
          </Typography>
          <Button sx={{ ml: "auto" }} size="sm" endDecorator={<AddIcon />}>
            Create
          </Button>
        </Box>
        <ListDivider sx={{ mt: 2 }} />
        <ListItem>
          <Avatar variant="outlined" color="primary" size="sm">
            MD
          </Avatar>
        </ListItem>
        <ListItem>
          <Avatar variant="outlined" color="success" size="sm">
            MD
          </Avatar>
        </ListItem>
        <ListItem>
          <Avatar variant="outlined" color="warning" size="sm">
            MD
          </Avatar>
        </ListItem>
      </List>
    </SheetBox>
  );
};

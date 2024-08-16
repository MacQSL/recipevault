import {
  Avatar,
  Button,
  List,
  ListDivider,
  ListItem,
  Typography,
} from "@mui/joy";
import { SheetBox } from "./SheetBox";

export const RecipeList = () => {
  return (
    <SheetBox
      sx={{
        height: "100%",
        width: 300,
      }}
    >
      <List>
        <ListItem>
          <Typography level="h4" component="h2">
            Recipes
          </Typography>
          <Button sx={{ ml: "auto" }}>Create</Button>
        </ListItem>
        <ListDivider sx={{ mt: 2 }} />
        <ListItem>
          <Avatar variant="outlined" color="primary">
            MD
          </Avatar>
        </ListItem>
        <ListItem>
          <Avatar variant="outlined" color="success">
            JD
          </Avatar>
        </ListItem>
        <ListItem>
          <Avatar variant="outlined" color="warning">
            HH
          </Avatar>
        </ListItem>
      </List>
    </SheetBox>
  );
};

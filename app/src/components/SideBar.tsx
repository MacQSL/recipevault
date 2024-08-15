import {
  List,
  ListDivider,
  ListItem,
  ListItemButton,
  ListSubheader,
  Sheet,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/query/cookbook-query";

export const SideBar = () => {
  const { data: cookbooks, isSuccess } = useQuery(getCookbooksWithRecipes());

  return (
    <Sheet
      variant="outlined"
      sx={{
        width: 250,
        mx: "auto",
        my: 4,
        borderRadius: "sm",
        boxShadow: "md",
      }}
    >
      <List size="lg">
        <ListSubheader>RecipeVault</ListSubheader>
        <ListItem>
          <ListItemButton>Home</ListItemButton>
        </ListItem>
        <ListDivider />
        <ListSubheader>Cookbooks</ListSubheader>
        {isSuccess
          ? cookbooks.map((cookbook) => (
              <ListItem>
                <ListItemButton>{cookbook.name}</ListItemButton>
              </ListItem>
            ))
          : null}
        <ListDivider />
        <ListItem>
          <ListItemButton>Settings</ListItemButton>
        </ListItem>
        <ListItem>
          <ListItemButton>Log out</ListItemButton>
        </ListItem>
      </List>
    </Sheet>
  );
};

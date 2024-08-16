import {
  List,
  ListDivider,
  ListItem,
  ListItemButton,
  ListSubheader,
  Skeleton,
  Typography,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/query/cookbook-query";
import { SheetBox } from "./SheetBox";

export const SideBar = () => {
  const { data: cookbooks, isLoading } = useQuery(getCookbooksWithRecipes());

  return (
    <SheetBox
      sx={{
        my: 2,
        ml: 2,
        width: 250,
      }}
    >
      <List size="md">
        <ListItem>
          <Typography level="h3" component="h1">
            RecipeVault
          </Typography>
        </ListItem>
        <ListItem sx={{ mt: 4, ".Mui-selected": { borderRadius: "lg" } }}>
          <ListItemButton>
            <ListItemButton>Dashboard</ListItemButton>
          </ListItemButton>
        </ListItem>
        <ListSubheader sx={{ fontSize: 12 }}>Cookbooks</ListSubheader>
        {cookbooks?.map((cookbook) => (
          <Skeleton loading={isLoading}>
            <ListItem key={cookbook.cookbook_id}>
              <ListItemButton>{cookbook.name}</ListItemButton>
            </ListItem>
          </Skeleton>
        ))}
        <ListDivider sx={{ mt: "auto" }} />
        <ListItem>
          <ListItemButton>Settings</ListItemButton>
        </ListItem>
        <ListItem>
          <ListItemButton>Log out</ListItemButton>
        </ListItem>
      </List>
    </SheetBox>
  );
};

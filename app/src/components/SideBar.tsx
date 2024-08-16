import {
  List,
  ListItem,
  ListItemButton,
  ListSubheader,
  Sheet,
  Skeleton,
  Typography,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/query/cookbook-query";

export const SideBar = () => {
  const { data: cookbooks, isLoading } = useQuery(getCookbooksWithRecipes());

  return (
    <Sheet
      variant="outlined"
      sx={{
        my: 2,
        ml: 2,
        p: 2,
        display: "flex",
        width: 250,
        borderRadius: "md",
      }}
    >
      <List size="md">
        <ListItem>
          <Typography level="h3" component="h1">
            RecipeVault
          </Typography>
        </ListItem>
        <ListItem sx={{ mt: 4 }}>
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
        <ListItem sx={{ mt: "auto" }}>
          <ListItemButton>Settings</ListItemButton>
        </ListItem>
        <ListItem>
          <ListItemButton>Log out</ListItemButton>
        </ListItem>
      </List>
    </Sheet>
  );
};

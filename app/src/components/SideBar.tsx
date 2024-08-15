import {
  List,
  ListDivider,
  ListItem,
  ListItemButton,
  ListSubheader,
  Sheet,
  Skeleton,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/query/cookbook-query";

export const SideBar = () => {
  const {
    data: cookbooks,
    isSuccess,
    isLoading,
  } = useQuery(getCookbooksWithRecipes());

  return (
    <Sheet
      variant="outlined"
      sx={{
        width: 250,
        borderRadius: "sm",
        boxShadow: "md",
      }}
    >
      <List size="lg">
        <ListSubheader sx={{ fontSize: 12 }}>RecipeVault</ListSubheader>
        <ListItem>
          <ListItemButton>Home</ListItemButton>
        </ListItem>
        <ListDivider />
        <ListSubheader sx={{ fontSize: 12 }}>Cookbooks</ListSubheader>
        {cookbooks?.map((cookbook) => (
          <Skeleton loading={isLoading}>
            <ListItem key={cookbook.cookbook_id}>
              <ListItemButton>{cookbook.name}</ListItemButton>
            </ListItem>
          </Skeleton>
        ))}
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

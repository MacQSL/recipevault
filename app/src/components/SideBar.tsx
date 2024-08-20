import {
  List,
  ListDivider,
  ListItem,
  ListItemButton,
  ListItemDecorator,
  ListSubheader,
  Typography,
} from "@mui/joy";
import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/query/cookbook-query";
import { SheetBox } from "./SheetBox";
import SpaceDashboardIcon from "@mui/icons-material/SpaceDashboard";
import SettingsIcon from "@mui/icons-material/Settings";
import LogoutIcon from "@mui/icons-material/Logout";
import AutoStoriesIcon from "@mui/icons-material/AutoStories";
import ExpandMoreIcon from "@mui/icons-material/ExpandMore";
import { useState } from "react";

export const SideBar = () => {
  const { data: cookbooks, isLoading } = useQuery(getCookbooksWithRecipes());

  const [showCookbooks, setShowCookbooks] = useState(false);

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

        <ListSubheader sx={{ mt: 4 }}>Favourites</ListSubheader>
        <ListItem>
          <ListItemButton>
            <ListItemDecorator>
              <SpaceDashboardIcon />
            </ListItemDecorator>
            Item
          </ListItemButton>
        </ListItem>

        <ListSubheader sx={{ mt: 4 }}>Menu</ListSubheader>
        <ListItem>
          <ListItemButton>
            <ListItemDecorator>
              <SpaceDashboardIcon />
            </ListItemDecorator>
            Dashboard
          </ListItemButton>
        </ListItem>

        <ListItem nested>
          <ListItemButton onClick={() => setShowCookbooks((s) => !s)}>
            <ListItemDecorator>
              <AutoStoriesIcon />
            </ListItemDecorator>
            Cookbooks
            <ListItemDecorator>
              <ExpandMoreIcon
                sx={{
                  ml: "auto",
                  transform: showCookbooks ? "rotate(180deg)" : "rotate(360)",
                  transition: showCookbooks
                    ? "transform 0.1s ease-in-out"
                    : "none",
                }}
              />
            </ListItemDecorator>
          </ListItemButton>

          <ListItem>
            <List marker="circle">
              {showCookbooks &&
                cookbooks?.map((cookbook) => (
                  <ListItem key={cookbook.cookbook_id}>
                    <ListItemButton>{cookbook.name}</ListItemButton>
                  </ListItem>
                ))}
            </List>
          </ListItem>
        </ListItem>

        <ListDivider sx={{ mt: "auto" }} />
        <ListItem>
          <ListItemButton>
            <ListItemDecorator>
              <SettingsIcon />
            </ListItemDecorator>
            Settings
          </ListItemButton>
        </ListItem>

        <ListItem>
          <ListItemButton>
            <ListItemDecorator>
              <LogoutIcon />
            </ListItemDecorator>
            Log out
          </ListItemButton>
        </ListItem>
      </List>
    </SheetBox>
  );
};

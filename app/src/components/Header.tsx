import { Avatar, Input, Sheet } from "@mui/joy";
import { ToggleMode } from "./ToggleMode";

export const Header = () => {
  return (
    <Sheet
      sx={{
        display: "flex",
        justifyContent: "space-between",
        p: 2,
        boxShadow: "md",
        borderRadius: "md",
      }}
    >
      <Input placeholder="Search for recipes..."></Input>
      <ToggleMode />
      <Avatar variant="outlined" color="primary">
        MD
      </Avatar>
    </Sheet>
  );
};

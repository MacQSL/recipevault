import { Avatar, Input } from "@mui/joy";
import { ToggleMode } from "./ToggleMode";
import { SheetBox } from "./SheetBox";

export const Header = () => {
  return (
    <SheetBox
      sx={{
        justifyContent: "space-between",
      }}
    >
      <Input placeholder="Search for recipes..." sx={{ width: 300 }}></Input>
      <ToggleMode />
      <Avatar variant="outlined" color="primary">
        MD
      </Avatar>
    </SheetBox>
  );
};

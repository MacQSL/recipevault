import { Avatar, Input } from "@mui/joy";
import { ToggleMode } from "./ToggleMode";
import { SheetBox } from "./SheetBox";
import SearchIcon from "@mui/icons-material/Search";

export const Header = () => {
  return (
    <SheetBox
      sx={{
        justifyContent: "space-between",
      }}
    >
      <Input
        placeholder="Search for recipes..."
        sx={{ width: 300 }}
        endDecorator={<SearchIcon />}
      ></Input>
      <ToggleMode />
      <Avatar variant="outlined" color="primary" size="sm">
        MD
      </Avatar>
    </SheetBox>
  );
};

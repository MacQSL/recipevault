import { Input, Sheet } from "@mui/joy";
import { ToggleMode } from "./ToggleMode";

export const Header = () => {
  return (
    <Sheet
      variant="outlined"
      sx={{
        display: "flex",
        justifyContent: "space-between",
        p: 2,
        borderRadius: "md",
      }}
    >
      <Input placeholder="Search for recipes..."></Input>
      <ToggleMode />
    </Sheet>
  );
};

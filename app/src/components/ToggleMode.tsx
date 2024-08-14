import { Button, useColorScheme } from "@mui/joy";

export const ToggleMode = () => {
  const { mode, setMode } = useColorScheme();

  return (
    <Button
      variant="outlined"
      onClick={() => {
        setMode(mode === "light" ? "dark" : "light");
      }}
    >
      toggle
    </Button>
  );
};

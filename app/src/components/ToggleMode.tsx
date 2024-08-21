import { Switch, useColorScheme } from "@mui/joy";

export const ToggleMode = () => {
  const { mode, setMode } = useColorScheme();

  return (
    <Switch
      checked={mode === "dark"}
      onChange={() => {
        setMode(mode === "light" ? "dark" : "light");
      }}
    />
  );
};

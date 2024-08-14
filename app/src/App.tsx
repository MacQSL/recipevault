import { CssVarsProvider } from "@mui/joy";
import "./App.css";
import { ToggleMode } from "./components/ToggleMode";

function App() {
  return (
    <CssVarsProvider>
      <ToggleMode />
    </CssVarsProvider>
  );
}

export default App;

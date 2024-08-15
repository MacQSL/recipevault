import { CssVarsProvider } from "@mui/joy";
import "./App.css";
// TODO: Drop app.css
import { ToggleMode } from "./components/ToggleMode";
import { Login } from "./components/Login";
import { SideBar } from "./components/SideBar";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <CssVarsProvider>
        <ToggleMode />
        <Login />
        <SideBar />
      </CssVarsProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;

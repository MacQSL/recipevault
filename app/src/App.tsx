import { CssVarsProvider, Sheet } from "@mui/joy";
import { Login } from "./components/Login";
import { SideBar } from "./components/SideBar";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Header } from "./components/Header";
import { RecipeCard } from "./components/RecipeCard";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <CssVarsProvider>
        <Header />
        <Sheet sx={{ p: 2, display: "flex", flexFlow: "column" }}>
          <SideBar />
          <Login />
          <RecipeCard
            recipe={{
              recipe_id: "a",
              cookbook_id: "b",
              url: null,
              name: "Banana Bread",
              description: "Family recipe banana bread",
            }}
          />
        </Sheet>
      </CssVarsProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;

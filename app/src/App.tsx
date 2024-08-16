import { Box, CssVarsProvider, Sheet } from "@mui/joy";
import { Login } from "./components/Login";
import { SideBar } from "./components/SideBar";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { RecipeCard } from "./components/RecipeCard";
import { Header } from "./components/Header";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <CssVarsProvider>
        <Sheet variant="soft" sx={{ display: "flex", minHeight: "100dvh" }}>
          <SideBar />
          <Box
            sx={{
              p: 2,
              flex: 1,
              display: "flex",
              flexDirection: "column",
              minWidth: 0,
              overflow: "auto",
            }}
          >
            <Header />
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
          </Box>
        </Sheet>
      </CssVarsProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;

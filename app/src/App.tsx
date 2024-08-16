import { Box, CssVarsProvider, Sheet } from "@mui/joy";
import { SideBar } from "./components/SideBar";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { Header } from "./components/Header";
import { RecipeList } from "./components/RecipeList";

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
              gap: 2,
              minWidth: 0,
              overflow: "auto",
            }}
          >
            <Header />
            <RecipeList />
          </Box>
        </Sheet>
      </CssVarsProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;

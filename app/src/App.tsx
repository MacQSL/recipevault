import { CssVarsProvider } from "@mui/joy";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";
import { RecipeList } from "./components/RecipeList";
import { RecipeView } from "./components/RecipeView";
import { AppLayout } from "./components/AppLayout";

const queryClient = new QueryClient();

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <CssVarsProvider>
        <AppLayout>
          <RecipeList />
          <RecipeView />
        </AppLayout>
      </CssVarsProvider>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  );
}

export default App;

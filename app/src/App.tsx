import { Footer } from "./components/Footer";
import { Header } from "./components/Header";

function App() {
  return (
    <div className="flex flex-col h-screen bg-light dark:bg-dark">
      <Header />
      <div className="flex-grow">main</div>
      <Footer />
    </div>
  );
}

export default App;

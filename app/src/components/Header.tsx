import { useState } from "react";

export const Header = () => {
  const [darkMode, setDarkMode] = useState(true);

  const toggleTheme = () => {
    setDarkMode((d) => !d);
    document.body.classList.toggle("dark");
  };

  return (
    <div className="flex justify-between">
      <h1 className="text-2xl font-bold text-dark dark:text-light">
        Recipe Vault
      </h1>
      <button onClick={toggleTheme}>{darkMode ? "light" : "dark"}</button>
    </div>
  );
};

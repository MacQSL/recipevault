import { Box, Sheet } from "@mui/joy";
import { Header } from "./Header";
import { SideBar } from "./SideBar";
import { PropsWithChildren } from "react";

export const AppLayout = (props: PropsWithChildren) => {
  return (
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
        <Box sx={{ display: "flex", gap: 2 }}>{props.children}</Box>
      </Box>
    </Sheet>
  );
};

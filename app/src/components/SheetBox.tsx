import { Sheet, SheetProps } from "@mui/joy";

export const SheetBox = (props: SheetProps) => {
  const { sx, ...rest } = props;
  return (
    <Sheet
      variant="outlined"
      sx={{ p: 2, display: "flex", boxShadow: "md", borderRadius: "md", ...sx }}
      {...rest}
    ></Sheet>
  );
};

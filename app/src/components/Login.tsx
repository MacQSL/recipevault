import {
  Button,
  FormControl,
  FormLabel,
  Input,
  Link,
  Sheet,
  Typography,
} from "@mui/joy";

export const Login = () => {
  return (
    <Sheet
      variant="outlined"
      sx={{
        width: 300,
        mx: "auto", // margin left & right
        my: 4,
        p: 3,
        display: "flex",
        flexDirection: "column",
        gap: 2,
        borderRadius: "sm",
        boxShadow: "md",
      }}
    >
      <div>
        <Typography level="h4" component="h1">
          Welcome to RecipeVault
        </Typography>
        <Typography>Sign in to continue.</Typography>
      </div>

      <FormControl>
        <FormLabel>Email</FormLabel>
        <Input name="email" type="email" placeholder="johndoe@email.com" />
      </FormControl>
      <FormControl>
        <FormLabel>Password</FormLabel>
        <Input name="password" type="password" placeholder="password" />
      </FormControl>

      <Button sx={{ mt: 1 }}>Log in</Button>
      <Typography
        endDecorator={<Link href="/sign-up">Sign up</Link>}
        fontSize="sm"
        sx={{ alignSelf: "center" }}
      >
        Don't have an account?
      </Typography>
    </Sheet>
  );
};

import { useQuery } from "@tanstack/react-query";
import { getCookbooksWithRecipes } from "../api/queries/cookbook-queries";

export const LandingPage = () => {
  const query = useQuery(getCookbooksWithRecipes());
  console.log(query.data);

  return <div>landing page</div>;
};

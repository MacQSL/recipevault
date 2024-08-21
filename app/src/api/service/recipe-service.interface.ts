export interface IRecipe {
  recipe_id: string;
  cookbook_id: string;
  url: string | null;
  name: string;
  description: string | null;
}

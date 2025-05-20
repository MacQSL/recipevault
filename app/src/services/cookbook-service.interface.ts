export interface ICookbookRecipes {
  cookbook_id: string;
  name: string;
  description: string | null;
  recipes: unknown[];
}

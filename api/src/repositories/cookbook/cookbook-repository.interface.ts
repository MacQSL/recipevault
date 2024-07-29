import { IRecipeRecord } from '../recipe/recipe-repository.interface.js';

export interface ICookbookRecord {
  cookbook_id: string;
  name: string;
  description: string | null;
}

export interface ICookbookWithRecipes extends ICookbookRecord {
  recipes: IRecipeRecord[];
}

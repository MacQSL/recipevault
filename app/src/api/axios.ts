import axios from "axios";

/**
 * RecipeVault Axios client
 *
 */
export const client = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

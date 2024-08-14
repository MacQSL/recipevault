import axios from "axios";

/**
 * RecipeVault Axios instance
 *
 */
export const instance = axios.create({
  baseURL: import.meta.env.VITE_API_URL,
});

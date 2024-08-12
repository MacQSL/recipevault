package models

// Cookbook omitting audit columns
type Cookbook struct {
	CookbookID  int     `json:"cookbook_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// Cookbook with Recipes omitting audit columns
type CookbookRecipes struct {
	Cookbook
	Recipes []Recipe `json:"recipes"`
}

// Recipe omitting audit columns
type Recipe struct {
	RecipeID    int     `json:"recipe_id"`
	CookbookID  int     `json:"cookbook_id"`
	Name        string  `json:"name"`
	Url         *string `json:"url"`
	Description *string `json:"description"`
}

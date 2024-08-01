package model

import "database/sql"

// Cookbook record omitting audit columns
type CookbookRecord struct {
	Cookbook_id int            `json:"cookbook_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

// Cookbook with associated []Recipes
type CookbookRecipes struct {
	Cookbook_id int            `json:"cookbook_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Recipes     []RecipeRecord `json:"recipes"`
}

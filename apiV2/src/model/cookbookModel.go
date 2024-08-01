package model

import "database/sql"

type CookbookRecord struct {
	Cookbook_id int
	Name        string
	Description sql.NullString
}

type CookbookRecipes struct {
	Cookbook_id int            `json:"cookbook_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Recipes     []RecipeRecord `json:"recipes"`
}

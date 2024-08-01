package model

// Recipe record omitting audit columns
type RecipeRecord struct {
	Recipe_id   int    `json:"recipe_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

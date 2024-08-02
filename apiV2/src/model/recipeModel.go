package model

// Recipe omitting audit columns
type Recipe struct {
	Recipe_id   int    `json:"recipe_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

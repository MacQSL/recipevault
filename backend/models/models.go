package models

// Cookbook omitting audit columns
type Cookbook struct {
	CookbookID  int     `json:"cookbook_id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

// Cookbook scan row ptrs
func (c *Cookbook) Ptrs() []any {
	return []any{&c.CookbookID, &c.Name, &c.Description}
}

// Recipe omitting audit columns
type Recipe struct {
	RecipeID    int     `json:"recipe_id"`
	CookbookID  int     `json:"cookbook_id"`
	Name        string  `json:"name"`
	Url         *string `json:"url"`
	Description *string `json:"description"`
}

// Recipe scan row ptrs
func (r *Recipe) Ptrs() []any {
	return []any{&r.RecipeID, &r.CookbookID, &r.Name, &r.Url, &r.Description}
}

// Ingredient omitting audit columns
type Ingredient struct {
	IngredientID int    `json:"ingredient_id"`
	RecipeID     int    `json:"recipe_id"`
	Measurement  string `json:"measurement"`
	Description  string `json:"description"`
	Lexorank     string `json:"lexorank"`
}

// Ingredient scan row ptrs
func (r *Ingredient) Ptrs() []any {
	return []any{&r.IngredientID, &r.RecipeID, &r.Measurement, &r.Description, &r.Lexorank}
}

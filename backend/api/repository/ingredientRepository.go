package repository

import (
	"database/sql"
	"recipevault/models"
)

type IngredientRepository struct {
	db *sql.DB
}

// NewIngredientRepository creates a new instance of IngredientRepository
func NewIngredientRepository(db *sql.DB) *IngredientRepository {
	return &IngredientRepository{
		db: db,
	}
}

// Get Ingredients by Recipe ID
func (r *IngredientRepository) GetIngredientsByRecipeID(recipeID int) ([]models.Ingredient, error) {
	query := `
    SELECT
      i.ingredient_id,
      i.recipe_id,
      i.measurement,
      i.description,
    FROM ingredient i
    WHERE i.recipe_id = $1`

	return FindRows[models.Ingredient](r.db, query, recipeID)
}

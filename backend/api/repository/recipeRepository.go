package repository

import (
	"database/sql"
	"recipevault/models"
)

type RecipeRepository struct {
	db *sql.DB
}

func NewRecipeRepository(db *sql.DB) *RecipeRepository {
	return &RecipeRepository{
		db: db,
	}
}

// Get Recipe by ID
func (r *RecipeRepository) GetRecipeByID(recipeID int) (models.Recipe, error) {
	query := `
    SELECT
      recipe_id,
      cookbook_id,
      name,
      url,
      description
    FROM recipe
    WHERE recipe_id = $1;`

	return GetRow[models.Recipe](r.db, query, recipeID)
}

// Get Recipes by Cookbook ID
func (r *RecipeRepository) GetRecipesByCookbookID(cookbookID int) ([]models.Recipe, error) {
	query := `
    SELECT
      r.recipe_id,
      r.cookbook_id,
      r.name,
      r.url,
      r.description
    FROM recipe r
    JOIN cookbook c
    ON r.cookbook_id = c.cookbook_id
    WHERE r.cookbook_id = $1;`

	return FindRows[models.Recipe](r.db, query, cookbookID)
}

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
func (r *RecipeRepository) GetRecipeByID(recipeID int) (*models.Recipe, error) {
	row := r.db.QueryRow(`
    SELECT
      recipe_id,
      name,
      url,
      description
    FROM recipe
    WHERE recipe_id = $1;`, recipeID)

	c := &models.Recipe{}

	err := row.Scan(&c.Recipe_id, &c.Name, &c.Url, &c.Description)

	if err != nil {
		return nil, err
	}

	return c, nil
}

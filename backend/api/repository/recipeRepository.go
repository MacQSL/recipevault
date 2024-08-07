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
      cookbook_id,
      name,
      url,
      description
    FROM recipe
    WHERE recipe_id = $1;`, recipeID)

	c := &models.Recipe{}

	err := row.Scan(&c.Recipe_id, &c.Cookbook_id, &c.Name, &c.Url, &c.Description)

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Get Recipes by Cookbook ID
func (r *RecipeRepository) GetRecipesByCookbookID(cookbookID int) ([]models.Recipe, error) {
	rows, err := r.db.Query(`
    SELECT
      r.recipe_id,
      r.cookbook_id,
      r.name,
      r.url,
      r.description
    FROM recipe r
    JOIN cookbook c
    ON r.cookbook_id = c.cookbook_id
    WHERE r.cookbook_id = $1;`, cookbookID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recipes := []models.Recipe{}

	for rows.Next() {
		c := models.Recipe{}

		err := rows.Scan(&c.Recipe_id, &c.Cookbook_id, &c.Name, &c.Url, &c.Description)

		if err != nil {
			return nil, err
		}

		recipes = append(recipes, c)
	}

	return recipes, nil
}

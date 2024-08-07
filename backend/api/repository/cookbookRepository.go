package repository

import (
	"database/sql"
	"encoding/json"
	"recipevault/models"
)

type CookbookRepository struct {
	db *sql.DB
}

func NewCookbookRepository(db *sql.DB) *CookbookRepository {
	return &CookbookRepository{
		db: db,
	}
}

// Get Cookbooks by User ID
func (r *CookbookRepository) GetCookbooksWithRecipesByUserID(userID int) ([]models.CookbookRecipes, error) {

	rows, err := r.db.Query(`
    SELECT
      c.cookbook_id,
      c.name,
      c.description,
      json_agg(json_build_object(
        'recipe_id', r.recipe_id,
        'cookbook_id', r.cookbook_id,
        'name', r.name,
        'url', r.url,
        'description', r.description
      )) as recipes
    FROM cookbook c
    INNER JOIN user_cookbook u
    ON c.cookbook_id = u.cookbook_id
    LEFT JOIN recipe r
    ON c.cookbook_id = r.cookbook_id
    WHERE u.user_id = $1
    GROUP BY c.cookbook_id;`, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cookbooks []models.CookbookRecipes

	for rows.Next() {
		var c models.CookbookRecipes
		var recipesJSON []byte

		err := rows.Scan(&c.Cookbook_id, &c.Name, &c.Description, &recipesJSON)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(recipesJSON, &c.Recipes)

		if err != nil {
			return nil, err
		}

		cookbooks = append(cookbooks, c)
	}

	return cookbooks, nil
}

// Get Cookbook by ID
func (r *CookbookRepository) GetCookbookByID(cookbookID int) (*models.Cookbook, error) {

	row := r.db.QueryRow(`
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    WHERE c.cookbook_id = $1;`, cookbookID)

	c := &models.Cookbook{}

	err := row.Scan(&c.Cookbook_id, &c.Name, &c.Description)

	if err != nil {
		return c, err
	}

	return c, nil
}

// Get Cookbook by cookbook ID and user ID
func (r *CookbookRepository) GetCookbookByIDAndUserID(cookbookID int, userID int) (*models.Cookbook, error) {

	row := r.db.QueryRow(`
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    INNER JOIN user_cookbook u
    ON c.cookbook_id = u.cookbook_id
    WHERE c.cookbook_id = $1
    AND u.user_id = $2
    ;`, cookbookID, userID)

	c := &models.Cookbook{}

	err := row.Scan(&c.Cookbook_id, &c.Name, &c.Description)

	if err != nil {
		return c, err
	}

	return c, nil
}

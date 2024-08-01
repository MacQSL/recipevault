package repository

import (
	"database/sql"
	"recipehub/api/src/model"
)

type CookbookRepository struct {
	db *sql.DB
}

func NewCookbookRepository(db *sql.DB) *CookbookRepository {
	return &CookbookRepository{
		db: db,
	}
}

// Get Cookbooks by User ID - Includes associated recipes
func (r *CookbookRepository) GetCookbooksByUserID(userID int) ([]model.CookbookRecipes, error) {

	rows, err := r.db.Query(`
    SELECT c.*
    FROM cookbook c
    JOIN user u
    ON c.user_id = ?;
    `, userID)

	if err != nil {
		return nil, err
	}

	cookbooks := []model.CookbookRecipes{}

	for rows.Next() {
		c := model.CookbookRecipes{}

		err := rows.Scan(&c.Cookbook_id, &c.Name, &c.Recipes, &c.Description)

		if err != nil {
			return nil, err
		}

		cookbooks = append(cookbooks, c)
	}

	return cookbooks, nil
}

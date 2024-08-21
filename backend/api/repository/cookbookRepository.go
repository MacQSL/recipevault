package repository

import (
	"database/sql"
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
func (r *CookbookRepository) GetCookbooksByUserID(userID int) ([]models.Cookbook, error) {
	query := `
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    INNER JOIN user_cookbook u
    ON c.cookbook_id = u.cookbook_id
    WHERE u.user_id = $1
    GROUP BY c.cookbook_id
    ORDER BY c.name;`

	return FindRows[models.Cookbook](r.db, query, userID)
}

// Get Cookbook by ID
func (r *CookbookRepository) GetCookbookByID(cookbookID int) (models.Cookbook, error) {
	query := `
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    WHERE c.cookbook_id = $1;`

	return GetRow[models.Cookbook](r.db, query, cookbookID)
}

// Get Cookbook by cookbook ID and user ID
func (r *CookbookRepository) GetCookbookByIDAndUserID(cookbookID int, userID int) (models.Cookbook, error) {
	query := `
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    INNER JOIN user_cookbook u
    ON c.cookbook_id = u.cookbook_id
    WHERE c.cookbook_id = $1
    AND u.user_id = $2 ;`

	return GetRow[models.Cookbook](r.db, query, cookbookID, userID)
}

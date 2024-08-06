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

// Scan rows into cookbooks
func scanCookbooks(rows *sql.Rows) ([]models.Cookbook, error) {
	defer rows.Close()

	cookbooks := []models.Cookbook{}
	for rows.Next() {
		c := models.Cookbook{}
		err := rows.Scan(&c.Cookbook_id, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		cookbooks = append(cookbooks, c)
	}

	return cookbooks, nil
}

// Get Cookbooks by User ID
func (r *CookbookRepository) GetCookbooksByUserID(userID int) ([]models.Cookbook, error) {

	rows, err := r.db.Query(`
    SELECT
      c.cookbook_id,
      c.name,
      c.description
    FROM cookbook c
    INNER JOIN user_cookbook u
    ON c.cookbook_id = u.cookbook_id
    WHERE u.user_id = $1;`, userID)

	if err != nil {
		return nil, err
	}

	data, err := scanCookbooks(rows)

	return data, err
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

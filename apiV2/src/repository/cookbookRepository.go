package repository

import (
	"database/sql"
	"recipehub/api/src/model"
	"recipehub/api/src/util"
)

type CookbookRepository struct {
	log util.ILogger
	db  *sql.DB
}

func NewCookbookRepository(log util.ILogger, db *sql.DB) *CookbookRepository {
	return &CookbookRepository{
		log: log,
		db:  db,
	}
}

// Scan rows into cookbooks
func scanCookbooks(rows *sql.Rows) ([]model.Cookbook, error) {
	defer rows.Close()

	cookbooks := []model.Cookbook{}
	for rows.Next() {
		c := model.Cookbook{}
		err := rows.Scan(&c.Cookbook_id, &c.Name, &c.Description)
		if err != nil {
			return nil, err
		}
		cookbooks = append(cookbooks, c)
	}

	return cookbooks, nil
}

// Get Cookbooks by User ID
func (r *CookbookRepository) GetCookbooksByUserID(userID int) ([]model.Cookbook, error) {

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

package service

import (
	"database/sql"
	"recipevault/api/repository"
	"recipevault/models"
)

type IngredientService struct {
	repository *repository.IngredientRepository
}

// Construct new instance of IngredientService
func NewIngredientService(r *repository.IngredientRepository) *IngredientService {
	return &IngredientService{
		repository: r,
	}
}

// Initialize IngredientService with dependencies
func InitIngredientService(db *sql.DB) *IngredientService {
	r := repository.NewIngredientRepository(db)
	return NewIngredientService(r)
}

// Get Recipe Ingredients
func (s *IngredientService) GetRecipeIngredients(recipeID int) ([]models.Ingredient, error) {
	return s.repository.GetIngredientsByRecipeID(recipeID)
}

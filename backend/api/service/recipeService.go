package service

import (
	"database/sql"
	"recipevault/api/repository"
	"recipevault/models"
)

type RecipeService struct {
	repository        *repository.RecipeRepository
	ingredientService *IngredientService
}

// Construct new instance of RecipeService
func NewRecipeService(r *repository.RecipeRepository, i *IngredientService) *RecipeService {
	return &RecipeService{
		repository:        r,
		ingredientService: i,
	}
}

// Initialize RecipeService with dependencies
func InitRecipeService(db *sql.DB) *RecipeService {
	r := repository.NewRecipeRepository(db)
	i := InitIngredientService(db)
	return NewRecipeService(r, i)
}

// Get recipe
func (s *RecipeService) GetRecipe(recipeID int) (*models.RecipeIngredients, error) {
	r, err := s.repository.GetRecipeByID(recipeID)

	if err != nil {
		return nil, err
	}

	i, err := s.ingredientService.GetRecipeIngredients(recipeID)

	if err != nil {
		return nil, err
	}

	recipe := &models.RecipeIngredients{
		Recipe:      r,
		Ingredients: i,
	}

	return recipe, nil
}

// Get cookbook recipes
func (s *RecipeService) GetCookbookRecipes(cookbookID int) ([]models.Recipe, error) {
	return s.repository.GetRecipesByCookbookID(cookbookID)
}

package service

import (
	"database/sql"
	"recipevault/api/repository"
	"recipevault/models"
)

type RecipeService struct {
	repository *repository.RecipeRepository
}

// Construct new instance of RecipeService
func NewRecipeService(r *repository.RecipeRepository) *RecipeService {
	return &RecipeService{
		repository: r,
	}
}

// Initialize RecipeService with dependencies
func InitRecipeService(db *sql.DB) *RecipeService {
	r := repository.NewRecipeRepository(db)
	return NewRecipeService(r)
}

// Get recipe
func (s *RecipeService) GetRecipe(recipeID int) (*models.Recipe, error) {
	return s.repository.GetRecipeByID(recipeID)
}

// Get cookbook recipes
func (s *RecipeService) GetCookbookRecipes(cookbookID int) (*[]models.Recipe, error) {
	return s.repository.GetRecipesByCookbookID(cookbookID)
}

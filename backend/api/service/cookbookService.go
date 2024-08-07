package service

import (
	"database/sql"
	"recipevault/api/repository"
	"recipevault/models"
)

type CookbookService struct {
	repository    *repository.CookbookRepository
	recipeService *RecipeService
}

// Construct new instance of CookbookService
func NewCookbookService(r *repository.CookbookRepository, rs *RecipeService) *CookbookService {
	return &CookbookService{
		repository:    r,
		recipeService: rs,
	}
}

// Initialize CookbookService with dependencies
func InitCookbookService(db *sql.DB) *CookbookService {
	r := repository.NewCookbookRepository(db)
	rs := InitRecipeService(db)
	return NewCookbookService(r, rs)
}

// Get cookbook
func (s *CookbookService) GetCookbook(cookbookID int) (*models.Cookbook, error) {
	return s.repository.GetCookbookByID(cookbookID)
}

// Get user cookbooks
func (s *CookbookService) GetUserCookbooksWithRecipes(userID int) ([]models.CookbookRecipes, error) {
	return s.repository.GetCookbooksWithRecipesByUserID(userID)
}

// Get user cookbook
func (s *CookbookService) GetUserCookbook(cookbookID int, userID int) (*models.Cookbook, error) {
	return s.repository.GetCookbookByIDAndUserID(cookbookID, userID)
}

// Can user access the cookbook
func (s *CookbookService) CanUserAccessCookbook(cookbookID int, userID int) bool {
	_, err := s.repository.GetCookbookByIDAndUserID(cookbookID, userID)
	return err == nil
}

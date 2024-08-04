package service

import (
	"recipevault/api/repository"
	"recipevault/models"
)

type CookbookService struct {
	repository *repository.CookbookRepository
}

func NewCookbookService(repo *repository.CookbookRepository) *CookbookService {
	return &CookbookService{
		repository: repo,
	}
}

// Get cookbook
func (s *CookbookService) GetCookbook(cookbookID int) (*models.Cookbook, error) {
	return s.repository.GetCookbookByID(cookbookID)
}

// Get user cookbooks
func (s *CookbookService) GetUserCookbooks(userID int) ([]models.Cookbook, error) {
	return s.repository.GetCookbooksByUserID(userID)
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

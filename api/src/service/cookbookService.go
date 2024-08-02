package service

import (
	"recipevault/api/src/model"
	"recipevault/api/src/repository"
)

type CookbookService struct {
	repository *repository.CookbookRepository
}

func NewCookbookService(repo *repository.CookbookRepository) *CookbookService {
	return &CookbookService{
		repository: repo,
	}
}

// Get user cookbooks
func (s *CookbookService) GetUserCookbooks(userID int) ([]model.Cookbook, error) {
	return s.repository.GetCookbooksByUserID(userID)
}

// Get cookbook
func (s *CookbookService) GetCookbook(cookbookID int) (model.Cookbook, error) {
	return s.repository.GetCookbookByID(cookbookID)
}

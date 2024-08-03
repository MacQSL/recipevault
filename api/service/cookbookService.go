package service

import (
	"errors"
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

// Get cookbook
func (s *CookbookService) GetCookbook(cookbookID int) (model.Cookbook, error) {
	return s.repository.GetCookbookByID(cookbookID)
}

// Get user cookbooks
func (s *CookbookService) GetUserCookbooks(userID int) ([]model.Cookbook, error) {
	return s.repository.GetCookbooksByUserID(userID)
}

// Get user cookbook
func (s *CookbookService) GetUserCookbook(cookbookID int, userID int) (model.Cookbook, error) {
	cookbooks, err := s.repository.GetCookbooksByUserID(userID)
	cookbook := model.Cookbook{}
	if err != nil {
		return cookbook, err
	}
	// find cookbook with matching cookbook_id
	for i := range cookbooks {
		if cookbooks[i].Cookbook_id == cookbookID {
			return cookbooks[i], nil
		}
	}
	return cookbook, errors.New("cookbook not found")
}

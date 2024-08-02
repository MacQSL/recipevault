package service

import "recipehub/api/src/model"

type ICookbookRepository interface {
	GetCookbooksByUserID(userID int) ([]model.Cookbook, error)
}

type CookbookService struct {
	repository ICookbookRepository
}

func NewCookbookService(repository ICookbookRepository) *CookbookService {
	return &CookbookService{
		repository: repository,
	}
}

func (s *CookbookService) GetUserCookbooks(userID int) ([]model.Cookbook, error) {
	return s.repository.GetCookbooksByUserID(userID)
}

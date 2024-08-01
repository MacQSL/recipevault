package service

import "recipehub/api/src/model"

type ICookbookRepository interface {
	GetCookbooksByUserID(userID int) ([]model.CookbookRecipes, error)
}

type CookbookService struct {
	repository ICookbookRepository
}

func NewCookbookService(repository ICookbookRepository) *CookbookService {
	return &CookbookService{
		repository: repository,
	}
}

func (s *CookbookService) GetUserCookbooks(userID int) ([]model.CookbookRecipes, error) {
	return s.repository.GetCookbooksByUserID(userID)
}

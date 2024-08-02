package service

import (
	"recipevault/api/src/model"
	"recipevault/api/src/repository"
	"recipevault/api/src/util"
)

type CookbookService struct {
	log        util.ILogger
	repository *repository.CookbookRepository
}

func NewCookbookService(log util.ILogger, repo *repository.CookbookRepository) *CookbookService {
	return &CookbookService{
		repository: repo,
	}
}

func (s *CookbookService) GetUserCookbooks(userID int) ([]model.Cookbook, error) {
	return s.repository.GetCookbooksByUserID(userID)
}

package service

import (
	"github.com/sakuyacatcat/scrape-line-bot/pkg/domain/model"
	"github.com/sakuyacatcat/scrape-line-bot/pkg/domain/repository"
)

type CoatService struct {
	repo repository.CoatRepository
}

func NewCoatService(repo repository.CoatRepository) *CoatService {
	return &CoatService{repo: repo}
}

func (s *CoatService) Search() ([]model.Coat, error) {
	result, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}

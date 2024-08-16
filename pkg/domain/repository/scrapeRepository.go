package repository

import "github.com/sakuyacatcat/scrape-line-bot/pkg/domain/model"

type scrapeRepository struct{}

func NewScrapeRepository() CoatRepository {
	return &scrapeRepository{}
}

func (r *scrapeRepository) FindAll() ([]model.Coat, error) {
	return []model.Coat{}, nil
}

package repository

import "github.com/sakuyacatcat/scrape-line-bot/pkg/domain/model"

type CoatRepository interface {
	FindAll() ([]model.Coat, error)
}

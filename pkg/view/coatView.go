package view

import "github.com/sakuyacatcat/scrape-line-bot/pkg/domain/model"

type CoatView struct{}

func NewCoatView() *CoatView {
	return &CoatView{}
}

func (v *CoatView) Render(coats []model.Coat) (string, error) {
	return "coat", nil
}

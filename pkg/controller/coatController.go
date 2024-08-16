package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/sakuyacatcat/scrape-line-bot/pkg/domain/service"
	"github.com/sakuyacatcat/scrape-line-bot/pkg/view"
)

type CoatController struct {
	service *service.CoatService
	view    *view.CoatView
}

func NewCoatController(service *service.CoatService, view *view.CoatView) *CoatController {
	return &CoatController{
		service: service,
		view:    view,
	}
}

func (c *CoatController) GetCoat(url string) (string, error) {
	result, err := scrapeWeb(url)
	if err != nil {
		log.Printf("failed to scrape web: %v", err)
		result = "failed to scrape web"
	}

	return result, nil
}

func scrapeWeb(url string) (string, error) {
	// HTTPリクエストを送信
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to fetch the URL: %s", url)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	// 例として、タイトルを取得
	title := doc.Find("title").Text()
	return title, nil
}

package infrastructure

import (
	"fmt"
	"sync"

	"github.com/playwright-community/playwright-go"
)

var (
	scraper *Scraper
	once    sync.Once
)

type Scraper struct {
	pw      *playwright.Playwright
	browser playwright.Browser
}

type Page struct {
	page playwright.Page
}

// シングルトンパターンにより、インスタンスは一度しか生成されない
func GetScraper() (*Scraper, error) {
	var err error
	once.Do(func() {
		scraper, err = newScraper()
	})
	return scraper, err
}

func newScraper() (*Scraper, error) {
	pw, err := playwright.Run()
	if err != nil {
		return nil, fmt.Errorf("could not start Playwright: %w", err)
	}
	browser, err := pw.Chromium.Launch()
	if err != nil {
		return nil, fmt.Errorf("could not launch browser: %w", err)
	}

	return &Scraper{
		pw:      pw,
		browser: browser,
	}, nil
}

func (sm *Scraper) NewPage() (*Page, error) {
	page, err := sm.browser.NewPage()
	if err != nil {
		return nil, fmt.Errorf("could not create page: %w", err)
	}

	return &Page{
		page: page,
	}, nil
}

func (sm *Scraper) Close() {
	sm.browser.Close()
	sm.pw.Stop()
}

func (s *Page) Close() {
	s.page.Close()
}

func (s *Page) Fetch(url string) error {
	_, err := s.page.Goto(url)
	if err != nil {
		return fmt.Errorf("could not navigate to the URL: %w", err)
	}
	return nil
}

func (s *Page) Extract(selector string) (string, error) {
	element := s.page.Locator(selector).First()
	if element == nil {
		return "", fmt.Errorf("element with selector: %s not found", selector)
	}
	text, err := element.TextContent()
	if err != nil {
		return "", fmt.Errorf("could not get text content: %w", err)
	}
	return text, nil
}

func (s *Page) Click(selector string) error {
	err := s.page.Locator(selector).Click()
	if err != nil {
		return fmt.Errorf("could not click element with selector: %s, error: %w", selector, err)
	}
	return nil
}

func (s *Page) Enter(selector, value string) error {
	err := s.page.Locator(selector).Fill(value)
	if err != nil {
		return fmt.Errorf("could not enter value into element with selector: %s, error: %w", selector, err)
	}
	return nil
}

func (s *Page) ExtractRow(selector string) (string, error) {
	element := s.page.Locator(selector).First()
	if element == nil {
		return "", fmt.Errorf("row with selector: %s not found", selector)
	}
	text, err := element.InnerText()
	if err != nil {
		return "", fmt.Errorf("could not get row text content: %w", err)
	}
	return text, nil
}

package scraper

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

type ScrapeError struct {
	Err     error
	Message string
	Context map[string]interface{}
	Scraper string
}

func (err ScrapeError) Error() string {
	return "[" + err.Scraper + "] " + err.Message + ": " + err.Error()
}

func NewScrapeError(scraper, message string, err error, context map[string]interface{}) error {
	return ScrapeError{
		Err:     err,
		Message: message,
		Context: context,
		Scraper: scraper,
	}
}

type ScrapeData struct {
	Index int
	URL   string
	Title string
	Data  *goquery.Selection
	Err   error
}

type Scraper interface {
	Scrape(ctx context.Context, url string) <-chan ScrapeData
}

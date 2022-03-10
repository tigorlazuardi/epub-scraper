package scraper

import (
	"context"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
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

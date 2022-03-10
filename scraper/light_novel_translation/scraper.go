package lightnoveltranslation

import (
	"context"

	"github.com/tigorlazuardi/epub-scraper/scraper"
)

var _ scraper.Scraper = (*LNTScraper)(nil)

type LNTScraper struct {
	bufferSize int
	client     scraper.Doer
}

func (lntscraper *LNTScraper) Scrape(ctx context.Context, url string) <-chan scraper.ScrapeData {
	panic("not implemented") // TODO: Implement
}

func (lntscraper LNTScraper) scrapeSite(ctx context.Context, url string, channel chan<- scraper.ScrapeData) {
}

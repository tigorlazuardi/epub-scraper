package lightnoveltranslation

import (
	"context"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/tigorlazuardi/epub-scraper/scraper"
)

var _ scraper.Scraper = (*LNTScraper)(nil)

type LNTScraper struct {
	bufferSize int
	client     scraper.Doer
	wg         *sync.WaitGroup
}

func (lntscraper *LNTScraper) Scrape(ctx context.Context, url string) <-chan scraper.ScrapeData {
	c := make(chan scraper.ScrapeData, lntscraper.bufferSize)

	go func() {
		lntscraper.wg.Add(1)

		lntscraper.wg.Wait()
		close(c)
	}()

	return c
}

func (lntscraper LNTScraper) scrapeSite(ctx context.Context, url string, index int, channel chan<- scraper.ScrapeData) {
	defer lntscraper.wg.Done()
	var err error
	defer func() {
		if err != nil {
			channel <- scraper.ScrapeData{Err: err}
		}
	}()
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	res, err := lntscraper.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {

	}

}

func (LNTScraper) findNext(sel *goquery.Selection) (url string) {
	if nodes := sel.Find("p.alignright").Nodes; len(nodes) > 0 {
		if a := nodes[0].FirstChild; a != nil {
			for _, attr := range a.Attr {
				if attr.Key == "href" {
					return attr.Val
				}
			}
		}
	}
	return
}

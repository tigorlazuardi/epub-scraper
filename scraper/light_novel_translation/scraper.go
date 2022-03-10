package lightnoveltranslation

import (
	"context"
	"errors"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/tigorlazuardi/epub-scraper/scraper"
)

var _ scraper.Scraper = (*LNTScraper)(nil)

const site = "Light Novel Translation"

type LNTScraper struct {
	bufferSize int
	client     scraper.Doer
	wg         *sync.WaitGroup
	semaphore  chan struct{}
}

func (lntscraper *LNTScraper) Scrape(ctx context.Context, url string) <-chan scraper.ScrapeData {
	c := make(chan scraper.ScrapeData, lntscraper.bufferSize)

	lntscraper.wg.Add(1)
	go func() {
		lntscraper.scrapeSite(ctx, url, 0, c)
		lntscraper.wg.Done()
		lntscraper.wg.Wait()
		close(c)
	}()

	return c
}

func (lntscraper LNTScraper) scrapeSite(ctx context.Context, url string, index int, channel chan<- scraper.ScrapeData) {
	const selector = "div.entry-content"
	lntscraper.semaphore <- struct{}{}
	defer lntscraper.wg.Done()
	var err error
	defer func() {
		<-lntscraper.semaphore
		if err != nil {
			channel <- scraper.ScrapeData{Err: err}
		}
	}()

	logCtx := map[string]interface{}{
		"url":      url,
		"index":    index,
		"selector": selector,
	}

	doc, err := lntscraper.fetchSite(ctx, url)
	if err != nil {
		return
	}

	if next := lntscraper.findNext(doc); next != "" {
		lntscraper.wg.Add(1)
		go lntscraper.scrapeSite(ctx, next, index+1, channel)
	}

	content := doc.Find(selector).Nodes
	if len(content) == 0 {
		msg := "current site does not have content"
		err = scraper.NewScrapeError(site, msg, errors.New(msg), logCtx)
		return
	}

	// TODO: Implement further
}

func (LNTScraper) findNext(doc *goquery.Document) (url string) {
	if nodes := doc.Find("p.alignright").Nodes; len(nodes) > 0 {
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

func (lntscraper LNTScraper) fetchSite(ctx context.Context, url string) (doc *goquery.Document, err error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	logCtx := map[string]interface{}{
		"url":    url,
		"method": http.MethodGet,
	}
	res, err := lntscraper.client.Do(req)
	if err != nil {
		msg := "failed to open connection to target website"
		return nil, scraper.NewScrapeError(site, msg, err, logCtx)
	}
	defer res.Body.Close()

	logCtx["status_code"] = res.StatusCode
	if res.StatusCode >= 400 {
		msg := "unexpected error code from server"
		return nil, scraper.NewScrapeError(site, msg, errors.New(msg), logCtx)
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = scraper.NewScrapeError(site, "failed to parse website to html", err, logCtx)
	}

	return
}

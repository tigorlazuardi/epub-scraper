package scraper

import (
	"context"
	"errors"
	"net/http"
	"runtime"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/tigorlazuardi/epub-scraper/logger"
	"golang.org/x/net/html"
)

type GenericScraper struct {
	queueSize       int
	client          Doer
	wg              *sync.WaitGroup
	semaphore       chan struct{}
	domain          string
	contentSelector string
	titleSelector   string
	nextSelector    string
	// 1 Based index digit.
	startIndex int
	// 1 Based index digit.
	endIndexReducer int
}

type ScraperOption struct {
	Domain          string
	ContentSelector string
	TitleSelector   string
	NextSelector    string
	// 1 Based index digit.
	StartIndex int
	// 1 Based index digit.
	EndIndexReducer int
	Threads         int
	Client          Doer
}

func NewGenericScraper(opts *ScraperOption) *GenericScraper {
	if opts.Threads == 0 {
		opts.Threads = runtime.NumCPU()
	}
	if opts.Client == nil {
		opts.Client = http.DefaultClient
	}
	return &GenericScraper{
		queueSize:       opts.Threads,
		client:          opts.Client,
		wg:              &sync.WaitGroup{},
		semaphore:       make(chan struct{}, opts.Threads),
		domain:          opts.Domain,
		contentSelector: opts.ContentSelector,
		titleSelector:   opts.TitleSelector,
		nextSelector:    opts.NextSelector,
		startIndex:      opts.StartIndex,
		endIndexReducer: opts.EndIndexReducer,
	}
}

func (generic *GenericScraper) Scrape(ctx context.Context, url string) <-chan ScrapeData {
	c := make(chan ScrapeData, generic.queueSize)

	go func() {
		// no need for lint because the consumer will wait until the channel is closed.
		generic.wg.Add(1) //nolint:staticcheck
		generic.scrapeSite(ctx, url, 0, c)
		generic.wg.Wait()
		close(c)
	}()

	return c
}

/*
get the final product from current site.

waitgroup must be added one point before calling this because this will call done on end.
*/
func (generic GenericScraper) scrapeSite(ctx context.Context, url string, index int, channel chan<- ScrapeData) {
	generic.semaphore <- struct{}{}
	defer generic.wg.Done()
	var err error
	defer func() {
		<-generic.semaphore
		if err != nil {
			channel <- ScrapeData{Err: err, URL: url, Index: index}
		}
	}()

	logCtx := map[string]interface{}{
		"url":               url,
		"index":             index,
		"content-selector":  generic.contentSelector,
		"title-selector":    generic.titleSelector,
		"next-selector":     generic.nextSelector,
		"start-index":       generic.startIndex,
		"end-index-reducer": generic.endIndexReducer,
	}

	doc, err := generic.fetchSite(ctx, url)
	if err != nil {
		msg := "failed to get site data"
		err = NewScrapeError(generic.domain, msg, err, logCtx)
		return
	}

	if next := generic.findNext(doc); next != "" {
		generic.wg.Add(1)
		go generic.scrapeSite(ctx, next, index+1, channel)
	}

	content := doc.Find(generic.contentSelector)
	nodes := content.Nodes
	if len(nodes) > generic.startIndex {
		nodes = nodes[generic.startIndex-1:]
	}
	if len(nodes)-generic.endIndexReducer <= 0 {
		msg := "current site does not have content"
		err = NewScrapeError(generic.domain, msg, errors.New(msg), logCtx)
		return
	}
	nodes = nodes[:len(nodes)-generic.endIndexReducer]
	logCtx["end-index"] = len(nodes)

	data := ScrapeData{
		URL:   url,
		Index: index,
		Data:  nodes,
	}

	if title := doc.Find(generic.titleSelector).First(); title != nil {
		data.Title = title.Text()
	}

	logCtx["scrape-data"] = data
	logger.Info("scrape-data-info", "data", logCtx)
	channel <- data
}

func (generic GenericScraper) findNext(doc *goquery.Document) (url string) {
	for _, node := range doc.Find(generic.nextSelector).Nodes {
		if url := generic.urlFinder(node); url != "" {
			return url
		}
	}
	return
}

// Gets the first href found under the node tree. Returns empty string if no href found.
func (generic GenericScraper) urlFinder(node *html.Node) (url string) {
	for _, node := range goquery.NewDocumentFromNode(node).Find("[href]").Nodes {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				return attr.Val
			}
		}
	}
	return
}

func (generic GenericScraper) fetchSite(ctx context.Context, url string) (doc *goquery.Document, err error) {
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	logCtx := map[string]interface{}{
		"url":    url,
		"method": http.MethodGet,
	}
	res, err := generic.client.Do(req)
	if err != nil {
		msg := "failed to open connection to target website"
		return nil, NewScrapeError(generic.domain, msg, err, logCtx)
	}
	defer res.Body.Close()

	logCtx["status_code"] = res.StatusCode
	if res.StatusCode >= 400 {
		msg := "unexpected status code from server"
		return nil, NewScrapeError(generic.domain, msg, errors.New(msg), logCtx)
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		err = NewScrapeError(generic.domain, "failed to parse website to html", err, logCtx)
	}

	return
}

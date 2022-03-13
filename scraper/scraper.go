package scraper

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/tigorlazuardi/epub-scraper/logger"
	"github.com/tigorlazuardi/epub-scraper/unsafeutils"
	"golang.org/x/net/html"
)

type Doer interface {
	Do(*http.Request) (*http.Response, error)
}

var _ logger.Display = (*ScrapeError)(nil)

type ScrapeError struct {
	Err     error
	Message string
	Context map[string]interface{}
	Scraper string
}

func (scrapeerror *ScrapeError) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"message": scrapeerror.Message,
		"context": scrapeerror.Context,
		"scraper": scrapeerror.Scraper,
	}

	if e, errMarshal := json.Marshal(scrapeerror.Err); errMarshal == nil {
		if unsafeutils.GetString(e) == "{}" {
			m["error"] = scrapeerror.Err.Error()
		} else {
			m["error"] = json.RawMessage(e)
		}
	} else {
		m["error"] = scrapeerror.Err.Error()
	}

	return json.Marshal(m)
}

func (err ScrapeError) Error() string {
	return "[" + err.Scraper + "] " + err.Message + ": " + err.Err.Error()
}

func (scrape ScrapeError) Display() json.RawMessage {
	// const indent = "    "
	errBytes, err := json.Marshal(scrape.Err)
	if err != nil || unsafeutils.GetString(errBytes) == "{}" {
		errBytes = unsafeutils.GetBytes(scrape.Error())
	}
	v := make(map[string]interface{}, 4)

	v["error"] = json.RawMessage(errBytes)
	v["message"] = scrape.Message
	v["scraper"] = scrape.Scraper

	logCtx := make(map[string]interface{}, len(scrape.Context))
	for key, val := range scrape.Context {
		if display, ok := val.(logger.Display); ok {
			logCtx[key] = display.Display()
		} else {
			logCtx[key] = val
		}
	}

	b, _ := json.Marshal(v)
	return b
}

func NewScrapeError(scraper, message string, err error, context map[string]interface{}) error {
	if err == nil {
		panic("scrapeerror: wrapped error must not be nil")
	}
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
	Data  []*html.Node
	Err   error
}

type Scraper interface {
	Scrape(ctx context.Context, url string) <-chan ScrapeData
}

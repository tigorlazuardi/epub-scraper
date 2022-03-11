package logger

import (
	"os"

	log "github.com/inconshreveable/log15"
)

const appName = "EPUB-SCRAPER"

var (
	info  = log.New("app", appName)
	warn  = log.New("app", appName)
	error = log.New("app", appName)
)

var (
	Info  = info.Info
	Warn  = warn.Warn
	Error = error.Error
)

func init() {
	info.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))
	error.SetHandler(log.StreamHandler(os.Stderr, log.JsonFormat()))
	warn.SetHandler(log.DiscardHandler())
}

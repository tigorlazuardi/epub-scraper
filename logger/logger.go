package logger

import (
	"os"

	"github.com/inconshreveable/log15"
)

const appName = "EPUB-SCRAPER"

var (
	info  = log15.New("app", appName)
	warn  = log15.New("app", appName)
	error = log15.New("app", appName)
	panic = log15.New("app", appName)
	Info  = info.Info
	Warn  = warn.Warn
	Error = error.Error
	Panic = panic.Crit
)

func init() {
	info.SetHandler(log15.StreamHandler(os.Stdout, log15.JsonFormat()))
	error.SetHandler(log15.StreamHandler(os.Stderr, log15.JsonFormat()))
	warn.SetHandler(log15.DiscardHandler())
	panic.SetHandler(log15.StreamHandler(os.Stderr, log15.JsonFormat()))
}

package logger

import (
	"encoding/json"
	"io"
	"os"

	log "github.com/inconshreveable/log15"
)

type Display interface {
	Display() json.RawMessage
}

type M = map[string]interface{}

var (
	infoLogger  = log.New("level", "info")
	warnLogger  = log.New("level", "warn")
	errorLogger = log.New("level", "error")
	debugLogger = log.New("level", "debug")
)

func Info(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	infoLogger.Info(msg, ctx)
}

func Warn(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	warnLogger.Warn(msg, ctx)
}

func Error(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	errorLogger.Error(msg, ctx)
}

func Debug(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	debugLogger.Debug(msg, ctx)
}

func UpdateLogger(level int, writer io.Writer, fmt log.Format) {
	infoLogger.SetHandler(log.DiscardHandler())
	errorLogger.SetHandler(log.DiscardHandler())
	warnLogger.SetHandler(log.DiscardHandler())
	debugLogger.SetHandler(log.DiscardHandler())
	switch {
	case level > 0:
		errorLogger.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 1:
		warnLogger.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 2:
		infoLogger.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 3:
		debugLogger.SetHandler(log.StreamHandler(writer, fmt))
	}
}

func init() {
	UpdateLogger(1, os.Stdout, log.JsonFormat())
}

func toDisplay(ctx ...interface{}) []interface{} {
	for i := range ctx {
		if display, ok := ctx[i].(Display); ok {
			ctx[i] = display.Display()
		}
	}
	return ctx
}

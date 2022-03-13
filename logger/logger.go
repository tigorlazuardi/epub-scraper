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

var (
	info  = log.New("level", "info")
	warn  = log.New("level", "warn")
	error = log.New("level", "error")
	debug = log.New("level", "debug")
)

func Info(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	info.Info(msg, ctx)
}

func Warn(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	warn.Warn(msg, ctx)
}

func Error(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	error.Error(msg, ctx)
}

func Debug(msg string, ctx ...interface{}) {
	ctx = toDisplay(ctx)
	debug.Debug(msg, ctx)
}

func UpdateLogger(level int, writer io.Writer, fmt log.Format) {
	info.SetHandler(log.DiscardHandler())
	error.SetHandler(log.DiscardHandler())
	warn.SetHandler(log.DiscardHandler())
	debug.SetHandler(log.DiscardHandler())
	switch {
	case level > 0:
		error.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 1:
		warn.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 2:
		info.SetHandler(log.StreamHandler(writer, fmt))
		fallthrough
	case level > 3:
		debug.SetHandler(log.StreamHandler(writer, fmt))
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

package logger

import (
	"io"
	"os"

	log "github.com/inconshreveable/log15"
)

var (
	info  = log.New("level", "info")
	warn  = log.New("level", "warn")
	error = log.New("level", "error")
	debug = log.New("level", "debug")
)

var (
	Info  = info.Info
	Warn  = warn.Warn
	Error = error.Error
	Debug = debug.Debug
)

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

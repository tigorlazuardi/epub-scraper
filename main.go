package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/tigorlazuardi/epub-scraper/cli"
	"github.com/tigorlazuardi/epub-scraper/logger"
)

func main() {
	ctx, drop := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for range c {
			drop()
		}
	}()
	err := cli.Execute(ctx)
	if err != nil {
		logger.Error("failed to execute command", "error", err)
		os.Exit(1)
	}
	os.Exit(0)
}

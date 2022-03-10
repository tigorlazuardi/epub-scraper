package cli

import (
	"context"

	"github.com/spf13/cobra"
)

var cliCMD = &cobra.Command{
	Use: "epub-scraper",
}

func Execute(ctx context.Context) error {
	return cliCMD.ExecuteContext(ctx)
}

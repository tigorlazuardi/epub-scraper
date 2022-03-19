package cli

import (
	"errors"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "scrapes given url",
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("0 arguments given for scrape command. Refer lnt-scraper scrape --help for example.")
		}

		// conf, err := config.EnsureRead()
		// if err != nil {
		// 	return err
		// }

		// for _, v := range args {
		// 	cfg, err := conf.GetDomainConfig(v)
		// 	if err != nil {
		//
		// 	}
		// }

		return nil
	},
	Example: "epub-scraper scrape https://domain.name.com/rest/of/the/path",
}

func init() {
	cliCMD.AddCommand(scrapeCmd)
}

// func scrape() {
//
// }

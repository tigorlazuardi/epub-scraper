package config

import (
	"net/url"

	"github.com/tigorlazuardi/epub-scraper/logger"
)

func (c *Config) AddConfig(uri string, cfg *DomainConfiguration) (err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return NewConfigError("failed to parse url", err, logger.M{
			"url": uri,
		})
	}

	host := u.Hostname()

	_ = host

	return
}

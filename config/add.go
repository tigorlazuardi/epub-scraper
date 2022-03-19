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

	c.Domain[host] = cfg

	err = Write(c, GetFileName())
	if err != nil {
		return NewConfigError("failed to save config to file when adding new domain", err, logger.M{
			"url":    uri,
			"domain": host,
			"config": cfg,
		})
	}

	return
}

package config

import (
	"net/url"

	"github.com/tigorlazuardi/epub-scraper/logger"
)

type Config struct {
	Timeout int     `yaml:"timeout"`
	Threads int     `yaml:"threads"`
	Domain  Domains `yaml:"domain"`
}

type DomainConfiguration struct {
	ContentSelector string `yaml:"content_selector"`
	TitleSelector   string `yaml:"title_selector"`
	NextSelector    string `yaml:"next_selector"`
	StartIndex      int    `yaml:"start_index"`
	EndIndexReducer int    `yaml:"end_index_reducer"`
}

type Domains map[string]DomainConfiguration

// Create new config.
func NewDefaultConfig() *Config {
	return &Config{
		Threads: 4,
		Domain:  Domains{},
	}
}

func (c Config) GetDomainConfig(uri string) (cfg *DomainConfiguration, err error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, NewConfigError("failed to parse url", err, logger.M{
			"url": uri,
		})
	}

	host := u.Hostname()
	x, ok := c.Domain[host]
	if !ok {
		msg, err := logger.NewError("website does not exist in configuration")
		return nil, NewConfigError(msg, err, logger.M{"hostname": u.Hostname()})
	}
	return &x, nil
}

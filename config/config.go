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
	// Always starts from 1 unless on edit operation
	StartIndex      int `yaml:"start_index"`
	EndIndexReducer int `yaml:"end_index_reducer"`
}

func (cfg *DomainConfiguration) UpdateFromOther(other *DomainConfiguration) {
	if other.ContentSelector != "" {
		cfg.ContentSelector = other.ContentSelector
	}
	if other.TitleSelector != "" {
		cfg.TitleSelector = other.TitleSelector
	}

	if other.NextSelector != "" {
		cfg.NextSelector = other.NextSelector
	}

	if other.StartIndex != 0 {
		cfg.StartIndex = other.StartIndex
	}

	if other.EndIndexReducer != 0 {
		cfg.EndIndexReducer = other.EndIndexReducer
	}
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
		return nil, NewConfigError(msg, err, logger.M{"web_domain": u.Hostname(), "url": uri})
	}
	return &x, nil
}

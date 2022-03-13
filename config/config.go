package config

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

func NewDefaultConfig() *Config {
	return &Config{
		Threads: 4,
		Domain:  Domains{},
	}
}

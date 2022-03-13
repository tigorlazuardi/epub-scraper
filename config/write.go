package config

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/tigorlazuardi/epub-scraper/logger"
	"gopkg.in/yaml.v3"
)

func Write(cfg *Config, path string) error {
	_, err := os.Stat(path)
	var file *os.File
	if errors.Is(err, fs.ErrNotExist) {
		file, err = os.Create(path)
	} else {
		file, err = os.Open(path)
	}
	if err != nil {
		return NewConfigError("failed to open/create config file", err, logger.M{"path": path})
	}
	defer file.Close()
	return newEncoder(file).Encode(cfg)
}

func newEncoder(w io.Writer) *yaml.Encoder {
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(configIndent)
	return encoder
}

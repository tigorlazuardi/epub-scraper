package config

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

func Write(cfg *Config, path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		defer file.Close()
		return newEncoder(file).Encode(cfg)
	} else {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		return newEncoder(file).Encode(cfg)
	}
}

func newEncoder(w io.Writer) *yaml.Encoder {
	encoder := yaml.NewEncoder(w)
	encoder.SetIndent(configIndent)
	return encoder
}

package config

import (
	"errors"
	"io/fs"
	"os"

	"github.com/kirsle/configdir"
	"github.com/tigorlazuardi/epub-scraper/logger"
	"gopkg.in/yaml.v3"
)

// Ensure directory is created.
func EnsureDirectory() error {
	configPath := configdir.LocalConfig(configDirName)
	err := configdir.MakePath(configPath)
	if err != nil {
		err = NewConfigError("failed to create directory for config files", err, logger.M{"path": configPath})
	}
	return err
}

/*
Ensure config file exists prior reading.

If config does not exist prior calling this function, default config is written to file and is returned.

Default config is returned on any error event.
*/
func EnsureRead() (*Config, error) {
	cfg := NewDefaultConfig()
	err := EnsureDirectory()
	if err != nil {
		return cfg, err
	}
	filePath := GetFileName()
	logCtx := logger.M{"path": filePath}

	_, err = os.Stat(filePath)
	if errors.Is(err, fs.ErrNotExist) {
		err = Write(cfg, filePath)
		return cfg, err
	} else if err != nil {
		return cfg, NewConfigError("failed to open file stat", err, logCtx)
	}
	f, err := os.Open(filePath)
	if err != nil {
		return cfg, NewConfigError("failed to open config file", err, logCtx)
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return cfg, NewConfigError("failed to parse config file as yaml", err, logCtx)
	}
	return cfg, nil
}

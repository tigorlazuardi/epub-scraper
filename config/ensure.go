package config

import (
	"errors"
	"io/fs"
	"os"

	"github.com/kirsle/configdir"
)

func EnsureDirectory() error {
	configPath := configdir.LocalConfig(configDirName)
	return configdir.MakePath(configPath)
}

/*
Ensure config file exists prior reading.

If config does not exist prior calling this function, default config is returned and default config is written to file.

Default config is returned on any error event.
*/
func EnsureRead() (*Config, error) {
	cfg := NewDefaultConfig()
	err := EnsureDirectory()
	if err != nil {
		return cfg, err
	}
	filePath := GetFileName()

	_, err = os.Stat(filePath)
	if errors.Is(err, fs.ErrNotExist) {
		file, err := os.Create(filePath)
		if err != nil {
			return cfg, err
		}
		defer file.Close()
	} else if err != nil {
		return cfg, err
	}

	return cfg, err
}

package config

import (
	"path/filepath"

	"github.com/kirsle/configdir"
)

func Exist() bool {
	return false
}

func GetDir() string {
	return configdir.LocalConfig(configDirName)
}

func GetFileName() string {
	return filepath.Join(GetDir(), configFilename)
}

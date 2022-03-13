package config

import (
	"encoding/json"

	"github.com/tigorlazuardi/epub-scraper/logger"
	"github.com/tigorlazuardi/epub-scraper/unsafeutils"
)

type ConfigError struct {
	Message string
	Err     error
	Context map[string]interface{}
}

func NewConfigError(message string, err error, ctx map[string]interface{}) error {
	return &ConfigError{
		Message: message,
		Err:     err,
		Context: ctx,
	}
}

func (configerror *ConfigError) Error() string {
	return configerror.Message + ": " + configerror.Err.Error()
}

func (config ConfigError) Display() json.RawMessage {
	m := map[string]interface{}{
		"message": config.Message,
		"error":   config.Err,
	}

	if display, ok := config.Err.(logger.Display); ok { //nolint
		m["error"] = display.Display()
	} else if errBytes, err := json.Marshal(config.Err); err != nil || unsafeutils.GetString(errBytes) == "{}" {
		m["error"] = config.Err.Error()
	}

	for k, v := range config.Context {
		if display, ok := v.(logger.Display); ok {
			config.Context[k] = display.Display()
		}
	}
	m["context"] = config.Context

	w, _ := json.Marshal(m)
	return json.RawMessage(w)
}

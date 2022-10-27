package log

import (
	"os"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/rs/zerolog"
)

func NewLogger(config *conf.Config) (*zerolog.Logger, error) {
	_, err := os.Open(config.Log.LogFile)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

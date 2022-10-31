package log

import (
	"os"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/rs/zerolog"
)

func NewLogger(config *conf.Config) (*zerolog.Logger, *os.File, error) {
	file, err := os.Create(config.Log.LogFile)
	if err != nil {
		return nil, nil, err
	}

	file, err = os.OpenFile(config.Log.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		return nil, nil, err
	}

	logger := zerolog.New(file).With().Timestamp().Logger()

	return &logger, file, nil
}

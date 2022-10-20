package api

import (
	"os"
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	"github.com/alekstet/linux_service_checker/maker"
	"github.com/alekstet/linux_service_checker/notifier"

	"github.com/rs/zerolog"
)

type store struct {
	Log       zerolog.Logger
	config    *conf.Config
	state     maker.ServicesInfo
	notifiers []notifier.Notifier
	maker     maker.Maker
	mutex     sync.Mutex
}

func NewStore(config *conf.Config) (*store, error) {
	makerImpl, err := maker.NewMaker(config)
	if err != nil {
		return nil, err
	}

	return &store{
		Log:       zerolog.New(os.Stdout).With().Timestamp().Logger(),
		config:    config,
		notifiers: []notifier.Notifier{},
		maker:     makerImpl,
		mutex:     sync.Mutex{},
	}, nil
}

package api

import (
	"os"
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	"github.com/alekstet/linux_service_checker/maker"
	"github.com/alekstet/linux_service_checker/notifier"

	"github.com/rs/zerolog"
)

type Store struct {
	config    *conf.Config
	state     maker.ServicesInfo
	Log       zerolog.Logger
	notifiers []notifier.Notifier
	maker     maker.Maker
	mutex     sync.Mutex
}

func NewStore(config *conf.Config) (*Store, error) {
	makerImpl, err := maker.NewMaker(config)
	if err != nil {
		return nil, err
	}

	return &Store{
		config:    config,
		Log:       zerolog.New(os.Stdout).With().Timestamp().Logger(),
		notifiers: []notifier.Notifier{},
		maker:     makerImpl,
		mutex:     sync.Mutex{},
	}, nil
}

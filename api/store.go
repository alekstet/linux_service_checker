package api

import (
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	"github.com/alekstet/linux_service_checker/maker"
	"github.com/alekstet/linux_service_checker/notifier"
)

type Store struct {
	config    *conf.Config
	mutex     sync.Mutex
	notifiers []notifier.Notifier
	maker     maker.Maker
}

func NewStore(config *conf.Config) (*Store, error) {
	makerStore, err := maker.NewStore(config)
	if err != nil {
		return nil, err
	}

	return &Store{
		config:    config,
		notifiers: []notifier.Notifier{},
		maker:     makerStore,
		mutex:     sync.Mutex{},
	}, nil
}

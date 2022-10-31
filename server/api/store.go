package api

import (
	"sync"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/alekstet/linux_service_checker/server/maker"
	"github.com/alekstet/linux_service_checker/server/notifier"

	"github.com/rs/zerolog"
)

type store struct {
	Logger    *zerolog.Logger
	state     maker.ServicesInfo
	notifiers []notifier.Notifier
	config    *conf.Config
	maker     maker.Maker
	mutex     sync.Mutex
}

func NewStore(maker maker.Maker, config *conf.Config, logger *zerolog.Logger) *store {
	return &store{
		Logger: logger,
		config: config,
		mutex:  sync.Mutex{},
		maker:  maker,
	}
}

package api

import (
	"os"
	"sync"

	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/alekstet/linux_service_checker/server/maker"
	"github.com/alekstet/linux_service_checker/server/notifier"

	"github.com/rs/zerolog"
)

type store struct {
	Log       zerolog.Logger
	state     maker.ServicesInfo
	notifiers []notifier.Notifier
	config    *conf.Config
	maker     maker.Maker
	mutex     sync.Mutex
}

func NewStore(maker maker.Maker, config *conf.Config) *store {
	return &store{
		Log:    zerolog.New(os.Stdout).With().Timestamp().Logger(),
		config: config,
		mutex:  sync.Mutex{},
		maker:  maker,
	}
}

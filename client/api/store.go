package api

import (
	"net/http"
	"os"
	"sync"

	"github.com/alekstet/linux_service_checker/client/conf"

	"github.com/rs/zerolog"
)

type store struct {
	Log    zerolog.Logger
	config *conf.Config
	client http.Client
	mutex  sync.Mutex
	state  map[string]ServiceInfo
}

func NewStore(config *conf.Config) *store {
	return &store{
		Log:    zerolog.New(os.Stdout).With().Timestamp().Logger(),
		config: config,
		client: http.Client{},
		mutex:  sync.Mutex{},
	}
}

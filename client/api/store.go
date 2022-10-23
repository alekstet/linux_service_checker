package api

import (
	"net/http"
	"os"

	"github.com/alekstet/linux_service_checker/client/conf"

	"github.com/rs/zerolog"
)

type store struct {
	Log    zerolog.Logger
	config *conf.Config
	client http.Client
}

func NewStore(config *conf.Config) *store {
	return &store{
		Log:    zerolog.New(os.Stdout).With().Timestamp().Logger(),
		config: config,
		client: http.Client{},
	}
}

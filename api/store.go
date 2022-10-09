package api

import (
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	"golang.org/x/crypto/ssh"
)

type Store struct {
	Config *conf.Config
	Client *ssh.Client
	State  servicesInfo
	M      sync.Mutex
}

func NewStore(config *conf.Config) *Store {
	return &Store{
		Config: config,
		M:      sync.Mutex{},
	}
}

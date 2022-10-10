package maker

import (
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	ssh2 "github.com/alekstet/linux_service_checker/ssh"
	"golang.org/x/crypto/ssh"
)

var _ Maker = (*Store)(nil)

type Store struct {
	config *conf.Config
	client *ssh.Client
	mutex  sync.Mutex
}

func NewStore(config *conf.Config) (*Store, error) {
	client, err := ssh2.GetClient(config)
	if err != nil {
		return nil, err
	}

	return &Store{
		config: config,
		mutex:  sync.Mutex{},
		client: client,
	}, nil
}

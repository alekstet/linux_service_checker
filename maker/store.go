package maker

import (
	"sync"

	"github.com/alekstet/linux_service_checker/conf"
	ssh2 "github.com/alekstet/linux_service_checker/ssh"
	"golang.org/x/crypto/ssh"
)

var _ Maker = (*MakerImpl)(nil)

type MakerImpl struct {
	config *conf.Config
	client *ssh.Client
	mutex  sync.Mutex
}

func NewMaker(config *conf.Config) (*MakerImpl, error) {
	client, err := ssh2.GetClient(config)
	if err != nil {
		return nil, err
	}

	return &MakerImpl{
		config: config,
		mutex:  sync.Mutex{},
		client: client,
	}, nil
}

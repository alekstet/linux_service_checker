package api

import (
	"sync"

	"github.com/alekstet/linux_service_checker/collector"
	"github.com/alekstet/linux_service_checker/conf"
	"github.com/alekstet/linux_service_checker/notifyer"
	"github.com/alekstet/linux_service_checker/notifyer/telegram"
	"golang.org/x/crypto/ssh"
)

type Store struct {
	Config    *conf.Config
	Client    *ssh.Client
	M         sync.Mutex
	Notifyer  notifyer.Notifyer
	Collector collector.Collector
}

func NewStore(config *conf.Config) *Store {
	return &Store{
		Config:    config,
		Notifyer:  telegram.NewTelegramClient(config.NotifyerPlatform.Telegram.Token),
		Collector: collector.NewStore(config),
		M:         sync.Mutex{},
	}
}

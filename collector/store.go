package collector

import "github.com/alekstet/linux_service_checker/conf"

var _ Collector = (*Store)(nil)

type Store struct {
	Config *conf.Config
}

func NewStore(config *conf.Config) *Store {
	return &Store{}
}

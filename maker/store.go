package maker

import "github.com/alekstet/linux_service_checker/conf"

var _ Maker = (*Store)(nil)

type Store struct {
	Config *conf.Config
}

func NewStore(config *conf.Config) *Store {
	return &Store{}
}

package collector

import (
	"sync"

	"github.com/alekstet/linux_service_checker/client/conf"
	"github.com/alekstet/linux_service_checker/client/notifier"
)

type collectorImpl struct {
	config    *conf.Config
	notifiers []notifier.Notifier
	mutex     sync.Mutex
	Collector
}

func NewCollector(config *conf.Config, collector Collector) *collectorImpl {
	impl := &collectorImpl{
		Collector: collector,
		config:    config,
		mutex:     sync.Mutex{},
	}

	return impl
}

package collector

import (
	"sync"
)

func (impl *collectorImpl) notify(change map[string][]ServiceInfo) {
	var wg sync.WaitGroup
	wg.Add(len(impl.notifiers) * len(change))
	for k, v := range change {
		for _, notifier := range impl.notifiers {
			notifier := notifier
			go func() {
				err := notifier.Notify(k, v[0].Active, v[1].Active, &wg)
				if err != nil {
					return
				}
			}()
		}
	}

	wg.Wait()
}

/* func (impl *collectorImpl) checkChange(servicesInfo ServicesInfo) map[string][]ServiceInfo {
	if impl.state == nil {
		return nil
	}

	result := make(map[string][]ServiceInfo)
	for k, v := range servicesInfo {
		if v.Active != impl.state[k].Active || v.Loaded != impl.state[k].Loaded {
			impl.mutex.Lock()
			result[k] = []ServiceInfo{v, impl.state[k]}
			impl.mutex.Unlock()
		}
	}
	return result
} */

func (impl *collectorImpl) prepareResponse(active string, servicesInfo *ServicesInfo) *ServicesInfo {
	result := make(ServicesInfo)
	if active == "all" {
		return servicesInfo
	}

	for name, serviceInfo := range *servicesInfo {
		switch active {
		case "active":
			if serviceInfo.Active == "active(exited)" || serviceInfo.Active == "active(running)" {
				result[name] = serviceInfo
			}
		case "inactive":
			if serviceInfo.Active == "inactive(dead)" {
				result[name] = serviceInfo
			}
		}
	}

	return &result
}

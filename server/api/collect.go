package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/alekstet/linux_service_checker/server/maker"
)

func (store *store) notify(change map[string][]maker.ServiceInfo) {
	var wg sync.WaitGroup
	wg.Add(len(store.notifiers) * len(change))
	for k, v := range change {
		for _, notifier := range store.notifiers {
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

func (store *store) checkChange(servicesInfo maker.ServicesInfo) map[string][]maker.ServiceInfo {
	if store.state == nil {
		return nil
	}

	result := make(map[string][]maker.ServiceInfo)
	for k, v := range servicesInfo {
		if v.Active != store.state[k].Active || v.Loaded != store.state[k].Loaded {
			store.mutex.Lock()
			result[k] = []maker.ServiceInfo{v, store.state[k]}
			store.mutex.Unlock()
		}
	}
	return result
}

func (store *store) Collect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	active := r.URL.Query().Get("show")

	servicesInfo, err := store.maker.Get(active)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	change := store.checkChange(*servicesInfo)
	if len(change) != 0 {
		fmt.Println("im here")
		store.notify(change)
	}

	store.state = *servicesInfo

	fmt.Println((*servicesInfo)["test"].Active)

	store.maker.Collect()

	fmt.Println((*servicesInfo)["test"].Active)

	jsonResp, err := json.Marshal(servicesInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/alekstet/linux_service_checker/maker"
)

func (store *Store) notify(change map[string][]maker.ServiceInfo) {
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

func (store *Store) changeChange(servicesInfo maker.ServicesInfo) map[string][]maker.ServiceInfo {
	result := make(map[string][]maker.ServiceInfo)
	for k, v := range servicesInfo {
		if v.Active != store.state[k].Active || v.Loaded != store.state[k].Loaded {
			fmt.Println("before:", v)
			fmt.Println("after:", store.state[k])
			store.mutex.Lock()
			result[k] = []maker.ServiceInfo{v, store.state[k]}
			store.mutex.Unlock()
		}
	}
	return result
}

func (store *Store) Collect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicesInfo := store.maker.Collect()

	change := store.changeChange(*servicesInfo)
	if len(change) != 0 {
		store.notify(change)
	}

	store.state = *servicesInfo

	jsonResp, err := json.Marshal(servicesInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

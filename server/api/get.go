package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/alekstet/linux_service_checker/server/maker"
)

func (store *store) GetOne(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	serviceInfo, err := store.maker.GetOne(name)
	if err != nil {
		store.Logger.Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured"))
		return
	}

	change := store.checkChange(name, *serviceInfo)
	if len(change) != 0 {
		result := make(map[string][]maker.ServiceInfo)
		result[name] = change
		store.notify(result)
	}

	servicesInfo := make(maker.ServicesInfo)
	servicesInfo[name] = *serviceInfo
	store.state = servicesInfo

	store.Write(w, servicesInfo)
}

func (store *store) Get(w http.ResponseWriter, r *http.Request) {
	active := r.URL.Query().Get("show")

	servicesInfo, err := store.maker.Get(active)
	if err != nil {
		store.Logger.Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		m := make(map[string]string)
		m["error"] = "error occured"
		resp, _ := json.Marshal(m)
		w.Write(resp)
		return
	}

	fmt.Println("after get", servicesInfo)

	change := store.checkChanges(*servicesInfo)
	fmt.Println("after ch", change, len(change))
	for _, v := range change {
		if len(v) != 0 {
			store.notify(change)
		}
	}

	store.state = *servicesInfo

	store.Write(w, *servicesInfo)
}

func (store *store) Write(w http.ResponseWriter, data maker.ServicesInfo) {
	jsonResp, err := json.Marshal(data)
	if err != nil {
		store.Logger.Err(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

func (store *store) notify(change map[string][]maker.ServiceInfo) {
	var wg sync.WaitGroup
	wg.Add(len(store.notifiers) * len(change))
	for service, status := range change {
		for _, notifier := range store.notifiers {
			notifier := notifier
			go func() {
				if len(status) != 0 {
					err := notifier.Notify(service, status[0].Active, status[1].Active, &wg)
					if err != nil {
						return
					}
				}
			}()
		}
	}

	wg.Wait()
}

func (store *store) checkChange(name string, info maker.ServiceInfo) []maker.ServiceInfo {
	if info.Active != store.state[name].Active || info.Loaded != store.state[name].Loaded {
		return []maker.ServiceInfo{info, store.state[name]}
	}

	return nil
}

func (store *store) checkChanges(servicesInfo maker.ServicesInfo) map[string][]maker.ServiceInfo {
	if store.state == nil {
		return nil
	}

	result := make(map[string][]maker.ServiceInfo)
	for name, info := range servicesInfo {
		change := store.checkChange(name, info)
		store.mutex.Lock()
		result[name] = change
		store.mutex.Unlock()
	}

	return result
}

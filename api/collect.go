package api

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

func (store *Store) change() {
	var wg sync.WaitGroup
	wg.Add(len(store.notifiers))
	for _, notifier := range store.notifiers {
		notifier := notifier
		go func() {
			err := notifier.Notify("", "", &wg)
			if err != nil {
				return
			}
		}()
	}

	wg.Wait()
}

func (store *Store) Collect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicesInfo := store.maker.Collect()

	store.change()

	jsonResp, err := json.Marshal(servicesInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

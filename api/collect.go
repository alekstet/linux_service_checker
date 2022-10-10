package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (store *Store) change() {
	for _, notifier := range store.notifiers {
		err := notifier.Notify("", "")
		if err != nil {
			return
		}
	}
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

package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (store *Store) Datas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicesInfo := store.Collector.Collect()
	/* isChanged := store.checkChange(*servicesInfo)
	fmt.Println(isChanged)
	if isChanged {
		store.Notifyer.Notify("key", "value")
	} */

	//store.State = *servicesInfo

	jsonResp, err := json.Marshal(servicesInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

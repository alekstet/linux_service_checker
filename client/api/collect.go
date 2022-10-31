package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var showQueryParams = "show"

func (store *store) Collect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	showParams := r.URL.Query().Get(showQueryParams)
	url := store.config.CollectorServer.ServerURL + store.config.CollectorServer.ServerPort
	urlPath := "/collect" + "?" + showQueryParams + "=" + showParams

	request, err := http.NewRequest("GET", url+urlPath, nil)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := store.client.Do(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(result)
}

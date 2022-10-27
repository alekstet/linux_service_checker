package api

import (
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

func (store *store) Make(w http.ResponseWriter, r *http.Request) {
	request, err := http.NewRequest("POST", store.config.CollectorServer.ServerURL+store.config.CollectorServer.ServerPort+"/make", r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	response, err := store.client.Do(request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	fmt.Println(response.StatusCode)
}

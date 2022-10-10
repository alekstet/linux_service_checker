package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Command string `json:"command"`
	Name    string `json:"name"`
}

func (store *Store) Make(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	var request Request
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = store.maker.Make(request.Command, request.Name)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

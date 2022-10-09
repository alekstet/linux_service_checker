package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

type Request struct {
	Command string `json:"command"`
	Name    string `json:"name"`
}

func (store *Store) Action(w http.ResponseWriter, r *http.Request) {
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

	cmd := exec.Command("sudo", "systemctl", request.Command, request.Name)
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

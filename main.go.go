package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/alekstet/linux_service_checker/conf"
)

type Req struct {
	Command string `json:"command"`
	Name    string `json:"name"`
}

type Data struct {
	ServicesNames []string
	ServerName    string
}

func (d *Data) Datas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	services_works := []string{}
	services_statuses := []string{}
	services_journals := []string{}
	services_infos := make(map[string][]string)

	for _, j := range d.ServicesNames {
		cmd_status := exec.Command("ssh", d.ServerName, "sudo", "systemctl", "status", j)
		stdout_status, err := cmd_status.CombinedOutput()
		if err != nil {
			w.WriteHeader(500)
		}

		cmd_journal := exec.Command("ssh", d.ServerName, "sudo", "journalctl", "-u", j, "-e", "-n")
		stdout_journal, err := cmd_journal.CombinedOutput()
		if err != nil {
			w.WriteHeader(500)
		}

		services_statuses = append(services_statuses, string(stdout_status))
		services_journals = append(services_journals, string(stdout_journal))
		words := strings.Split(string(stdout_status), " ")

		for _, word := range words {
			contain_run := strings.Contains(word, "(running)")
			contain_exit := strings.Contains(word, "(exited)")
			contain_dead := strings.Contains(word, "(dead)")
			contain_rest := strings.Contains(word, "(auto-restart)")
			if contain_run {
				services_works = append(services_works, "running")
			}
			if contain_exit {
				services_works = append(services_works, "exited")
			}
			if contain_dead {
				services_works = append(services_works, "dead")
			}
			if contain_rest {
				services_works = append(services_works, "auto-restart")
			}
		}
	}

	for i, j := range d.ServicesNames {
		info := []string{services_statuses[i], services_works[i], services_journals[i]}
		services_infos[j] = info
	}

	jsonResp, err := json.Marshal(services_infos)
	if err != nil {
		w.WriteHeader(500)
	}
	w.Write(jsonResp)
}

func Action(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
	}
	var req Req
	json.Unmarshal(body, &req)
	cmd := exec.Command("sudo", "systemctl", req.Command, req.Name)
	_, err = cmd.CombinedOutput()
	if err != nil {
		w.WriteHeader(500)
	}
}

func main() {
	services_names, server_url, server_name := conf.ReadConfig()
	d := Data{services_names, server_name}

	html := http.FileServer(http.Dir("./dist"))

	http.HandleFunc("/datas", d.Datas)
	http.HandleFunc("/action", Action)
	http.Handle("/", html)

	err := http.ListenAndServe(server_url, nil)
	if err != nil {
		log.Fatal(err)
	}
}

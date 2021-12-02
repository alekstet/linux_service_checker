package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/olebedev/config"
)

type test_struct struct {
	Command string `json:"command"`
	Name    string `json:"name"`
}

func main() {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	yamlString := string(file)

	cfg, err := config.ParseYaml(yamlString)
	if err != nil {
		panic(err)
	}

	services_names_, err := cfg.List("services_names")
	if err != nil {
		log.Panicf("Error with config file: %v", err)
	}

	server_url, err := cfg.String("server_url")
	if err != nil {
		panic(err)
	}
	log.Printf("Config read OK\n")

	html := http.FileServer(http.Dir("./dist"))

	datas := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")

		services_names := []string{}
		services_works := []string{}
		services_statuses := []string{}
		services_journals := []string{}
		services_infos := make(map[string][]string)

		for _, name := range services_names_ {
			switch name := name.(type) {
			case string:
				services_names = append(services_names, name)
			default:
				fmt.Printf("wrong type! %v %T\n", name, name)
			}
		}

		for _, j := range services_names {
			cmd := exec.Command("ssh", "at01@x14.se.molti.tech", "sudo", "systemctl", "status", j)
			//cmd := exec.Command("sudo", "systemctl", "status", j)
			stdout, _ := cmd.CombinedOutput()

			cmd1 := exec.Command("ssh", "at01@x14.se.molti.tech", "sudo", "journalctl", "-u", j, "-e", "-n")
			//cmd1 := exec.Command("sudo", "journalctl", "-u", j, "-e", "-n")
			stdout1, _ := cmd1.CombinedOutput()

			status := string(stdout)
			journal := string(stdout1)

			services_statuses = append(services_statuses, status)
			services_journals = append(services_journals, journal)
			words := strings.Split(status, " ")

			for _, word := range words {
				contain := strings.Contains(word, "(running)")
				contain1 := strings.Contains(word, "(exited)")
				contain2 := strings.Contains(word, "(dead)")
				contain3 := strings.Contains(word, "(auto-restart)")
				if contain {
					services_works = append(services_works, "running")
				}
				if contain1 {
					services_works = append(services_works, "exited")
				}
				if contain2 {
					services_works = append(services_works, "dead")
				}
				if contain3 {
					services_works = append(services_works, "auto-restart")
				}
			}
		}

		for i, j := range services_names {
			info := []string{services_statuses[i], services_works[i], services_journals[i]}
			services_infos[j] = info
		}

		jsonResp, err := json.Marshal(services_infos)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonResp)
	}

	action := func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var m test_struct
		json.Unmarshal(body, &m)
		cmd := exec.Command("sudo", "systemctl", m.Command, m.Name)
		_, err = cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
	}

	http.HandleFunc("/datas", datas)
	http.HandleFunc("/action", action)
	http.Handle("/", html)

	err = http.ListenAndServe(server_url, nil)
	if err != nil {
		log.Fatal(err)
	}
}

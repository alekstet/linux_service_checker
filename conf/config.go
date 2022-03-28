package conf

import (
	"io/ioutil"
	"log"

	"github.com/olebedev/config"
)

var services_names []string

func ReadConfig() ([]string, string, string) {
	file, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	yamlString := string(file)

	cfg, err := config.ParseYaml(yamlString)
	if err != nil {
		log.Fatal(err)
	}

	services_names_, err := cfg.List("services_names")
	if err != nil {
		log.Fatalf("Error with config file: %v", err)
	}

	for _, name := range services_names_ {
		switch name := name.(type) {
		case string:
			services_names = append(services_names, name)
		default:
			log.Fatalf("wrong type! %v %T\n", name, name)
		}
	}

	server_url, err := cfg.String("server_url")
	if err != nil {
		log.Fatal(err)
	}
	server_name, err := cfg.String("server_name")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Config read OK\n")

	return services_names, server_url, server_name
}

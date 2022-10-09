package cmd

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/alekstet/linux_service_checker/api"
	"github.com/alekstet/linux_service_checker/conf"
)

func Run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := conf.ReadConfig(*configPath)
	if err != nil {
		return fmt.Errorf("error while reading config %w", err)
	}

	store := api.NewStore(config)

	err = store.TestSSHconnection()
	if err != nil {
		return fmt.Errorf("error while tesing ssh connection %w", err)
	}

	api.InitRouter(store)

	err = http.ListenAndServe(config.ExecutionServer.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

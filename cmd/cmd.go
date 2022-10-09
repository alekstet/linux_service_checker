package cmd

import (
	"net/http"

	"github.com/alekstet/linux_service_checker/api"
	"github.com/alekstet/linux_service_checker/conf"
)

func Run() error {
	config, err := conf.ReadConfig()
	if err != nil {
		return err
	}

	store := api.NewStore(config)

	err = store.TestSSHconnection()
	if err != nil {
		return err
	}

	api.InitRouter(store)

	err = http.ListenAndServe(config.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

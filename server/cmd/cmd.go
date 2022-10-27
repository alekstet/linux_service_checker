package cmd

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/alekstet/linux_service_checker/server/api"
	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/alekstet/linux_service_checker/server/maker"
)

func Run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := conf.ReadConfig(*configPath)
	if err != nil {
		return fmt.Errorf("error while reading config: %w", err)
	}

	makerImpl, err := maker.NewMaker(config)
	if err != nil {
		return fmt.Errorf("error while creating maker: %w", err)
	}

	store := api.NewStore(makerImpl, config)
	api.InitRouter(store)
	err = store.RegisterNotifier()
	if err != nil {
		return fmt.Errorf("error while registering notifier: %w", err)
	}

	store.Log.Info().Msg("Server is starting...")

	err = http.ListenAndServe(config.ExecutionServer.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

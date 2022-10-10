package cmd

import (
	"flag"
	"fmt"
	"log"
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

	store, err := api.NewStore(config)
	if err != nil {
		return fmt.Errorf("error while creating store %w", err)
	}

	err = store.RegisterNotifier()
	if err != nil {
		log.Println(err)
	}

	api.InitRouter(store)

	store.Log.Info().Msg("Program is starting...")

	err = http.ListenAndServe(config.ExecutionServer.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

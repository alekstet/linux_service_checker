package cmd

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/alekstet/linux_service_checker/client/api"
	"github.com/alekstet/linux_service_checker/client/collector"
	"github.com/alekstet/linux_service_checker/client/conf"
)

func Run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := conf.ReadConfig(*configPath)
	if err != nil {
		return fmt.Errorf("error while reading config %w", err)
	}

	pool, err := collector.NewPGXPool(config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("error while creating pgx pool %w", err)
	}

	collectorImpl := collector.NewCollector(config, pool)
	if err != nil {
		return fmt.Errorf("error while creating collector %w", err)
	}

	err = collectorImpl.RegisterNotifier()
	if err != nil {
		return fmt.Errorf("error while register notifiers %w", err)
	}

	store := api.NewStore(config)
	api.InitRouter(store)

	store.Log.Info().Msg("Client is starting...")

	err = http.ListenAndServe(config.ExecutionServer.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

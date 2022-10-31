package cmd

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/alekstet/linux_service_checker/server/api"
	"github.com/alekstet/linux_service_checker/server/conf"
	"github.com/alekstet/linux_service_checker/server/db"
	"github.com/alekstet/linux_service_checker/server/log"
	"github.com/alekstet/linux_service_checker/server/maker"
	"github.com/alekstet/linux_service_checker/server/ssh"
)

func Run() error {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	config, err := conf.ReadConfig(*configPath)
	if err != nil {
		return fmt.Errorf("error while reading config: %w", err)
	}

	client, err := ssh.GetClient(config)
	if err != nil {
		return fmt.Errorf("error while getting client: %w", err)
	}

	connectionPool, err := db.GetDBConnectionPool(config.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("error while getting connection pool: %w", err)
	}

	logger, file, err := log.NewLogger(config)
	if err != nil {
		return fmt.Errorf("error while creating logger: %w", err)
	}

	defer file.Close()

	maker := maker.NewMaker(config, connectionPool, client)
	maker.Collect(config.ExecutionServer.PollPeriod)

	store := api.NewStore(maker, config, logger)
	api.InitRouter(store)
	err = store.RegisterNotifier()
	if err != nil {
		return fmt.Errorf("error while registering notifier: %w", err)
	}

	store.Logger.Info().Msg("Server is starting...")

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			maker.TruncateTable()
			store.Logger.Info().Msg("Server is finished...")
			os.Exit(1)
		case syscall.SIGTERM:
			maker.TruncateTable()
			store.Logger.Info().Msg("Server is finished...")
			os.Exit(1)
		}
	}()

	err = http.ListenAndServe(config.ExecutionServer.ServerPort, nil)
	if err != nil {
		return err
	}

	return nil
}

package maker

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

func (impl *makerImpl) Collect(period time.Duration) {
	go func() {
		for {
			impl.collect()
			time.Sleep(period)
		}
	}()
}

func (impl *makerImpl) collect() {
	if impl.checkEmptyTable() {
		impl.setTable()
	}

	var wg sync.WaitGroup
	wg.Add(len(impl.config.MonitoringServer.ServicesNames))

	for _, service := range impl.config.MonitoringServer.ServicesNames {
		go func(service string) {
			impl.collectError = impl.collectService(service, &wg)
			fmt.Println(impl.collectError)
		}(service)
	}

	fmt.Println("here")

	wg.Wait()
}

func (impl *makerImpl) collectService(service string, wg *sync.WaitGroup) error {
	defer wg.Done()

	fmt.Println("collectService")

	info, err := impl.getServiceInfo(service)
	if err != nil {
		return err
	}

	info.Journal, err = impl.getServiceJournal(service)
	if err != nil {
		return err
	}

	pool, err := impl.dbPool.Acquire(context.Background())
	if err != nil {
		return err
	}

	defer pool.Release()

	updateStatement := `UPDATE services SET description=$1, loaded=$2, active=$3, journal=$4 WHERE name=$5`
	_, err = pool.Exec(context.Background(), updateStatement, info.Description, info.Loaded, info.Active, info.Journal, strings.TrimSuffix(info.Name, ".service"))
	if err != nil {
		return err
	}

	return nil
}

func (store *makerImpl) getServiceJournal(serviceName string) (string, error) {
	output, err := store.getCommandOutput("journalctl" + " " + "-u" + " " + serviceName + " " + "-n")
	if err != nil {
		return "", fmt.Errorf("error while get service journal for service: %s, err: %w", serviceName, err)
	}

	return output, nil
}

func (store *makerImpl) getServiceInfo(serviceName string) (*ServiceInfo, error) {
	output, err := store.getCommandOutput("systemctl" + " " + "status" + " " + serviceName)
	if err != nil {
		return nil, fmt.Errorf("error while get service info for service: %s, err: %w", serviceName, err)
	}

	return store.serverInfoParser(output)
}

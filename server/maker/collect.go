package maker

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/jackc/pgx/v5"
)

var _ Maker = (*makerImpl)(nil)

func (impl *makerImpl) getRows(active string) (pgx.Rows, error) {
	var rows pgx.Rows

	pool, err := impl.dbPool.Acquire(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer pool.Release()

	switch active {
	case "all":
		updateStatement := `SELECT * FROM services`
		rows, err = pool.Query(context.Background(), updateStatement)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case "active":
		updateStatement := `SELECT * FROM services WHERE active=$1 OR active=$2`
		rows, err = pool.Query(context.Background(), updateStatement, "active(exited)", "active(running)")
		if err != nil {
			log.Println(err)
			return nil, err
		}
	case "inactive":
		updateStatement := `SELECT * FROM services WHERE active=$1`
		rows, err = pool.Query(context.Background(), updateStatement, "inactive(dead)")
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return rows, nil
}

func (impl *makerImpl) Get(active string) (*ServicesInfo, error) {
	rows, err := impl.getRows(active)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	servicesInfo := make(ServicesInfo)

	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			log.Println(err)
			return nil, err
		}

		var serviceInfo ServiceInfo

		serviceInfo.Name = values[0].(string)
		serviceInfo.Description = values[1].(string)
		serviceInfo.Loaded = values[2].(string)
		serviceInfo.Active = values[3].(string)
		serviceInfo.Journal = values[4].(string)

		servicesInfo[serviceInfo.Name] = serviceInfo

	}

	return &servicesInfo, nil
}

func (impl *makerImpl) Collect() *ServicesInfo {
	if impl.checkEmptyTable() {
		impl.setTable()
	}

	servicesInfo := make(ServicesInfo)

	var wg sync.WaitGroup
	wg.Add(len(impl.config.MonitoringServer.ServicesNames))

	for _, service := range impl.config.MonitoringServer.ServicesNames {
		go func(service string) {
			defer wg.Done()

			info, err := impl.getServiceInfo(service)
			if err != nil {
				log.Println(err)
				return
			}

			info.Journal, err = impl.getServiceJournal(service)
			if err != nil {
				log.Println(err)
				return
			}

			pool, err := impl.dbPool.Acquire(context.Background())
			if err != nil {
				log.Println(err)
				return
			}

			defer pool.Release()

			updateStatement := `UPDATE services SET description=$1, loaded=$2, active=$3, journal=$4 WHERE name=$5`
			_, err = pool.Exec(context.Background(), updateStatement, info.Description, info.Loaded, info.Active, info.Journal, strings.TrimSuffix(info.Name, ".service"))
			if err != nil {
				log.Println(err)
				return
			}

			impl.mutex.Lock()
			defer impl.mutex.Unlock()
			servicesInfo[info.Name] = *info
		}(service)
	}

	wg.Wait()

	return &servicesInfo
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

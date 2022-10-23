package maker

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func (impl *makerImpl) CollectProcess() {
	go func() {
		for {
			impl.collect()
			time.Sleep(time.Second * 2)
		}
	}()
}

func (impl *makerImpl) collect() *ServicesInfo {
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

func (store *makerImpl) serverInfoParser(output string) (*ServiceInfo, error) {
	splittedOutput := strings.Split(output, "\n")

	var firstLine, secondLine, thirdLine string

	firstLine = splittedOutput[0]
	secondLine = splittedOutput[1]
	thirdLine = splittedOutput[2]

	splittedFirstLine := strings.Split(firstLine, " ")
	name := splittedFirstLine[1]

	res := strings.Index(firstLine, "-")
	description := firstLine[res+2:]

	splittedSecondLine := strings.Split(secondLine, " ")
	var loadedStatus []string
	for _, word := range splittedSecondLine {
		if word != "" {
			loadedStatus = append(loadedStatus, word)
		}
	}

	loaded := loadedStatus[1]

	splittedThirdLine := strings.Split(thirdLine, " ")
	var activeStatus []string
	for _, word := range splittedThirdLine {
		if word != "" {
			activeStatus = append(activeStatus, word)
		}
	}

	active := activeStatus[1] + activeStatus[2]

	return &ServiceInfo{
		Name:        name,
		Description: description,
		Loaded:      loaded,
		Active:      active,
	}, nil
}

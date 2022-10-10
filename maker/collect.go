package maker

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

type serviceInfo struct {
	Name        string
	Description string
	Loaded      string
	Active      string
}

type servicesInfo map[string]serviceInfo

func (store *Store) Collect() *servicesInfo {
	servicesInfo := make(servicesInfo)

	var wg sync.WaitGroup
	wg.Add(len(store.config.MonitoringServer.ServicesNames))

	for _, service := range store.config.MonitoringServer.ServicesNames {
		go func(service string) {
			info, err := store.getServiceInfo(service, &wg)
			if err != nil {
				log.Println(err)
				return
			}

			_, err = store.getServiceJournal(service)
			if err != nil {
				log.Println(err)
				return
			}

			time.Sleep(time.Second * 2)

			store.mutex.Lock()
			defer store.mutex.Unlock()
			servicesInfo[info.Name] = *info
		}(service)
	}

	wg.Wait()

	return &servicesInfo
}

func (store *Store) getServiceJournal(serviceName string) (*servicesInfo, error) {
	output, err := store.getCommandOutput("journalctl" + " " + "-u" + serviceName + "-e" + "-n")
	if err != nil {
		return nil, err
	}

	splittedOutput := strings.Split(output, "\n")

	fmt.Println(splittedOutput)

	return nil, nil
}

func (store *Store) getCommandOutput(cmd string) (string, error) {
	var stdoutBuf bytes.Buffer

	session, err := store.client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	session.Stdout = &stdoutBuf
	session.Run(cmd)

	return stdoutBuf.String(), nil
}

func (store *Store) getServiceInfo(serviceName string, wg *sync.WaitGroup) (*serviceInfo, error) {
	defer wg.Done()

	output, err := store.getCommandOutput("systemctl" + " " + "status" + " " + serviceName)
	if err != nil {
		return nil, err
	}

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
		if word == "" {
			continue
		}

		loadedStatus = append(loadedStatus, word)
	}

	loaded := loadedStatus[1]

	vthirdLine := strings.Split(thirdLine, " ")
	var activeStatus []string
	for _, word := range vthirdLine {
		if word == "" {
			continue
		}

		activeStatus = append(activeStatus, word)
	}

	active := activeStatus[1] + activeStatus[2]

	return &serviceInfo{
		Name:        name,
		Description: description,
		Loaded:      loaded,
		Active:      active,
	}, nil
}

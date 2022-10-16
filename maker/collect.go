package maker

import (
	"log"
	"strings"
	"sync"
)

func (impl *MakerImpl) Collect() *ServicesInfo {
	servicesInfo := make(ServicesInfo)

	var wg sync.WaitGroup
	wg.Add(len(impl.config.MonitoringServer.ServicesNames) * 2)

	for _, service := range impl.config.MonitoringServer.ServicesNames {
		go func(service string) {
			info, err := impl.getServiceInfo(service, &wg)
			if err != nil {
				log.Println(err)
				return
			}

			info.Journal, err = impl.getServiceJournal(service, &wg)
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

func (store *MakerImpl) getServiceJournal(serviceName string, wg *sync.WaitGroup) (string, error) {
	defer wg.Done()

	output, err := store.getCommandOutput("journalctl" + " " + "-u" + " " + serviceName + " " + "-n")
	if err != nil {
		return "", err
	}

	return output, nil
}

func (store *MakerImpl) getServiceInfo(serviceName string, wg *sync.WaitGroup) (*ServiceInfo, error) {
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

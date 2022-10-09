package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"sync"
)

type serviceInfo struct {
	Name        string
	Description string
	Loaded      string
	Active      string
}

type servicesInfo map[string]serviceInfo

func (store *Store) getServiceJournal(serviceName string) (*servicesInfo, error) {
	return nil, nil
}

func (store *Store) getCommandOutput(serviceName string) (string, error) {
	var stdoutBuf bytes.Buffer

	session, err := store.Client.NewSession()
	if err != nil {
		return "", err
	}

	defer session.Close()

	session.Stdout = &stdoutBuf
	session.Run("systemctl" + " " + "status" + " " + serviceName)

	return stdoutBuf.String(), nil
}

func (store *Store) checkChange(servicesInfo servicesInfo) bool {
	return !reflect.DeepEqual(store.State, servicesInfo)
}

func (store *Store) getServiceInfo(serviceName string, wg *sync.WaitGroup) (*serviceInfo, error) {
	defer wg.Done()

	output, err := store.getCommandOutput(serviceName)
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

func (store *Store) getServicesInfo() *servicesInfo {
	servicesInfo := make(servicesInfo)

	var wg sync.WaitGroup
	wg.Add(len(store.Config.ServicesNames))

	for _, service := range store.Config.ServicesNames {
		go func(service string) {
			info, err := store.getServiceInfo(service, &wg)
			if err != nil {
				log.Println(err)
				return
			}

			store.M.Lock()
			defer store.M.Unlock()
			servicesInfo[info.Name] = *info
		}(service)
	}

	wg.Wait()

	return &servicesInfo
}

func (store *Store) notify() {

}

func (store *Store) Datas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	servicesInfo := store.getServicesInfo()
	isChanged := store.checkChange(*servicesInfo)
	fmt.Println(isChanged)
	if isChanged {
		store.notify()
	}

	store.State = *servicesInfo

	jsonResp, err := json.Marshal(servicesInfo)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)
}

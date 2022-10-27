package maker

import "strings"

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

package maker

import "fmt"

func (impl *makerImpl) Make(service, action string) error {
	_, err := impl.getCommandOutput("systemctl" + " " + action + " " + service)
	if err != nil {
		return fmt.Errorf("error while make command for service: %s, err: %w", service, err)
	}

	fmt.Println("all is ok")

	return nil
}

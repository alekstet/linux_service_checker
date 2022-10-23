package maker

import "fmt"

func (impl *makerImpl) Make(service, action string) error {
	fmt.Println(service, action)
	sss, err := impl.getCommandOutput("systemctl" + " " + action + " " + service)
	fmt.Println(sss, err)
	if err != nil {
		return fmt.Errorf("error while make command for service: %s, err: %w", service, err)
	}

	return nil
}

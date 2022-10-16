package maker

func (impl *MakerImpl) Make(service, action string) error {
	_, err := impl.getCommandOutput("systemctl" + " " + action + " " + service)
	if err != nil {
		return err
	}

	return nil
}

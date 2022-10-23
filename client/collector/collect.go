package collector

func (impl *collectorImpl) GetData() (*ServicesInfo, error) {
	_, err := impl.Collect()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

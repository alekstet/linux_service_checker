package collector

type Collector interface {
	Collect() *servicesInfo
}

package maker

type Maker interface {
	Collect() *servicesInfo
	Make(service, action string) error
}

package maker

type Maker interface {
	Make(service, action string) error
	Get(active string) (*ServicesInfo, error)
	GetOne(name string) (*ServiceInfo, error)
}

package maker

type Maker interface {
	Make(service, action string) error
	Collect() *ServicesInfo
	Get(active string) (*ServicesInfo, error)
}
